package websockit_test

import (
	"log"
	"time"

	"github.com/lossdev/websockit"
)

func ExampleWebsocket_ReadLoop() {
	// you should have a *websockit.WebsocketServer / *websockit.WebsocketClient already initialized
	var ws websockit.WebsocketClient

	readChan := make(chan []byte)
	go func() {
		if err := ws.ReadLoop(readChan); err != nil {
			log.Println(err)
		}
	}()

	for msg := range readChan {
		log.Printf("read: %s\n", string(msg))
	}
}

func ExampleWebsocketClient_ServerPingLoop() {
	// you should have a *websockit.WebsocketServer / *websockit.WebsocketClient already initialized
	var ws websockit.WebsocketClient

	pingOpts := []websockit.ClientPingOption{
		websockit.PingWithPongTimeout(10 * time.Second),
		websockit.PingWithPongLog(true),
	}
	ws.EnableServerPings(pingOpts...)

	go func() {
		if err := ws.ServerPingLoop(); err != nil {
			log.Println(err)
		}
	}()
}
