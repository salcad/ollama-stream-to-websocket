package service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"ollama-stream-to-websocket/src/main/config"
	"ollama-stream-to-websocket/src/main/lib"
	"ollama-stream-to-websocket/src/main/model"
)

const endOfTransmission = "\u0004" // End-of-Transmission character

// Stream LLM response to the WebSocket client
func StreamToWebSocket(ws *websocket.Conn, requestData model.RequestData, client *http.Client) {
	jsonStr, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshalling request data:", err)
		return
	}

	// Create the POST request using the URL from config
	req, err := http.NewRequest("POST", config.AppConfig.OllamaAPIUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Use the provided HTTP client to send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Stream the response to the WebSocket client
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var apiResponse model.APIResponse
		err := json.Unmarshal(scanner.Bytes(), &apiResponse)
		if err != nil {
			fmt.Println("Error unmarshalling response:", err)
			return
		}

		// Only send the "response" property to the WebSocket client
		err = ws.WriteMessage(websocket.TextMessage, []byte(apiResponse.Response))
		if err != nil {
			fmt.Println("Error writing message:", err)
			return
		}

		if apiResponse.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, []byte(endOfTransmission))
	if err != nil {
		fmt.Println("Error writing message:", err)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	var requestData model.RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	httpClient := &http.Client{}

	lib.ClientsLock.Lock()
	defer lib.ClientsLock.Unlock()

	for client := range lib.Clients {
		go StreamToWebSocket(client, requestData, httpClient)
	}
	fmt.Fprintf(w, "POST request received and being processed")
}

