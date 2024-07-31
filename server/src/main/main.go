package main

import (
	"fmt"
	"net/http"
	
	"ollama-stream-to-websocket/src/main/config"
	"ollama-stream-to-websocket/src/main/controller"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	mux := http.NewServeMux()
	controller.SetupRoutes(mux)

	fmt.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", config.EnableCORS(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
