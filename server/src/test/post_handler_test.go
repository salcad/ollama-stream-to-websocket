package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ollama-stream-to-websocket/src/main/model"
	"ollama-stream-to-websocket/src/main/service"
)

func TestPostHandler(t *testing.T) {
	// Given: Prepare the request data and environment
	requestData := model.RequestData{
		Model:  "llama3.1",
		Prompt: "Why is there something rather than nothing?",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		t.Fatalf("Failed to marshal requestData: %v", err)
	}

	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.PostHandler)

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
}

