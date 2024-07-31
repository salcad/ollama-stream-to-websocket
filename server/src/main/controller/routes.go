package controller

import (
	"net/http"
	"ollama-stream-to-websocket/src/main/service"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/post", service.PostHandler)
	mux.HandleFunc("/ws", service.WSHandler)
}
