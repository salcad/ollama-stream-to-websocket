package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"ollama-stream-to-websocket/src/main/lib"
	"ollama-stream-to-websocket/src/main/model"
	"ollama-stream-to-websocket/src/main/service"
)

const mockAPIResponse = `{"response":"Tokyo","done":false}
{"response":"New York City","done":false}
{"response":"Shanghai","done":true}`
type MockRoundTripper struct{}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(mockAPIResponse))),
		Header:     make(http.Header),
	}, nil
}

func TestPostHandlerWebSocket(t *testing.T) {
	// Given: Prepare the request data and mock WebSocket connection
	requestData := model.RequestData{
		Model:  "llama3.1",
		Prompt: "list 3 the busiest city in the world, just write only the name of the city, dont explain anything",
	}

	// Mock the HTTP client to return a predefined response
	mockHTTPClient := &http.Client{
		Transport: &MockRoundTripper{},
	}

	// Create a mock WebSocket server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade WebSocket: %v", err)
		}
		defer conn.Close()

		lib.ClientsLock.Lock()
		lib.Clients[conn] = true
		lib.ClientsLock.Unlock()

		// Wait until all messages are sent before closing the connection
		done := make(chan bool)
		go func() {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					done <- true
					break
				}
			}
		}()

		<-done
	}))
	defer server.Close()

	// Create a WebSocket client to connect to the mock server
	wsURL := "ws" + server.URL[len("http"):]
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to dial WebSocket: %v", err)
	}
	defer conn.Close()

	// Override the PostHandler to inject the mock HTTP client
	postHandler := func(w http.ResponseWriter, r *http.Request) {
		var requestData model.RequestData
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		lib.ClientsLock.Lock()
		defer lib.ClientsLock.Unlock()

		for client := range lib.Clients {
			go service.StreamToWebSocket(client, requestData, mockHTTPClient)
		}
		fmt.Fprintf(w, "POST request received and being processed")
	}

	// Create a new POST request with the JSON body
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		t.Fatalf("Failed to marshal requestData: %v", err)
	}
	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)

	// When: Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Then: Check the status code and response body
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "POST request received and being processed"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Verify the WebSocket messages received by the client
	expectedMessages := []string{"Tokyo", "New York City", "Shanghai", "\u0004"}
	for _, expectedMessage := range expectedMessages {
		_, message, err := conn.ReadMessage()
		if err != nil {
			t.Fatalf("Error reading WebSocket message: %v", err)
		}
		if string(message) != expectedMessage {
			t.Errorf("Expected message %s, but got %s", expectedMessage, message)
		}
	}
}
