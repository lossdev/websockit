// simple echo server using github.com/lossdev/websockit

package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/lossdev/websockit"
)

func main() {
	var addr = flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	ws := websockit.NewWebsocket()
	opts := []websockit.WebsocketServerOption{
		websockit.ServerWithHandshakeTimeout(5 * time.Second),
	}

	conn, err := ws.ServerSocket(w, req, nil, opts...)
	if err != nil {
		http.Error(w, "error upgrading to websocket connection", http.StatusInternalServerError)
		return
	}
	defer conn.CloseNice()
	readChan := make(chan []byte)

	go func() {
		if err := conn.ReadLoop(readChan); err != nil {
			log.Println(err)
		}
	}()

	for msg := range readChan {
		log.Printf("read msg: %s | client: %s\n", string(msg), conn.RemoteAddr().String())
		if err := conn.WriteTextMessage(msg); err != nil {
			log.Println(err)
		}
	}
}
