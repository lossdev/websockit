package websockit_test

import (
	"log"
	"net/http"
	"time"

	"github.com/lossdev/websockit"
)

func ExampleWebsocket_ServerSocket() {
	// this code should be placed in an HTTP handler function
	// i.e func foo(w http.ResponseWriter, r *http.Request) { ... }
	ws := websockit.NewWebsocket()
	opts := []websockit.WebsocketServerOption{
		websockit.ServerWithHandshakeTimeout(5 * time.Second),
	}

	// this is just so the snippet compiles - you'd already have `w` and `req` in scope and defined here
	var w http.ResponseWriter
	var req *http.Request

	if err := ws.ServerSocket(w, req, nil, opts...); err != nil {
		log.Println(err)
	}
}

func ExampleWebsocket_ClientSocket() {
	ws := websockit.NewWebsocket()
	opts := []websockit.WebsocketClientOption{
		websockit.ClientWithHandshakeTimeout(5 * time.Second),
		websockit.ClientWithProxy(http.ProxyFromEnvironment),
	}
	if err := ws.ClientSocket("wss://ws.postman-echo.com/raw", nil, opts...); err != nil {
		log.Println(err)
	}
}
