package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"ollama-stream-to-websocket/src/main/lib"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()

	lib.ClientsLock.Lock()
	lib.Clients[ws] = true
	lib.ClientsLock.Unlock()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			lib.ClientsLock.Lock()
			delete(lib.Clients, ws)
			lib.ClientsLock.Unlock()
			break
		}
	}
}
