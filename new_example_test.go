package websockit_test

import (
	"log"
	"net/http"
	"time"

	"github.com/lossdev/websockit"
)

func ExampleWebsocket_ServerSocket() {
	ws := websockit.NewWebsocket()
	opts := []websockit.WebsocketOption{
		websockit.WithHandshakeTimeout(5 * time.Second),
		websockit.WithProxy(http.ProxyFromEnvironment),
	}
	if err := ws.ServerSocket("http://localhost:8000", nil, opts...); err != nil {
		log.Println(err)
	}
}

func ExampleWebsocket_ClientSocket() {
	ws := websockit.NewWebsocket()
	opts := []websockit.WebsocketOption{
		websockit.WithHandshakeTimeout(5 * time.Second),
		websockit.WithProxy(http.ProxyFromEnvironment),
	}
	if err := ws.ClientSocket("http://localhost:8000", nil, opts...); err != nil {
		log.Println(err)
	}
}
