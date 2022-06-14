package websockit

import (
	"github.com/gorilla/websocket"
)

// Websocket is the embedded struct that websocket clients and servers have in common
type Websocket struct {
	conn   *websocket.Conn
	dialer *websocket.Dialer
}

// WebsocketServer inherits the Websocket fields and methods, and implements its own unique methods
type WebsocketServer struct {
	Websocket
}

// WebsocketClient inherits the Websocket fields and methods, and implements its own unique methods
type WebsocketClient struct {
	Websocket
}
