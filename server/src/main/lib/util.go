package lib

import (
	"sync"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)
var ClientsLock sync.Mutex
