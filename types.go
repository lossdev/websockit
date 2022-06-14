package websockit

import (
	"github.com/gorilla/websocket"
)

type Websocket struct {
	Conn   *websocket.Conn
	Dialer *websocket.Dialer
}

type WebsocketServer struct {
	Websocket
}

type WebsocketClient struct {
	Websocket
}
