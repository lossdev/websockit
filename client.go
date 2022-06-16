package websockit

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type pingOpts struct {
	pongTimeout time.Duration
	logPongs    bool
}

// ClientPingOption specifies any custom options for the client ping behavior
type ClientPingOption func(*pingOpts)

// ErrPingNotEnabled is returned when a client attempts to enter a ping loop without initializing the ping options
var ErrPingNotEnabled = errors.New("ping functionality has not been enabled yet; make a call to *WebsocketClient.EnableServerPings() before *WebsocketClient.ServerPingLoop()")

// EnableServerPings initializes the variables controlling the options needed to enter the ping loop, and should be called before ServerPingLoop
func (w *WebsocketClient) EnableServerPings(opts ...ClientPingOption) {
	pongTimeout := 60 * time.Second
	po := &pingOpts{pongTimeout, false}
	for _, o := range opts {
		o(po)
	}
	w.pongTimeout = po.pongTimeout
	_ = w.conn.SetReadDeadline(time.Now().Add(po.pongTimeout))
	if po.logPongs {
		w.conn.SetPongHandler(func(string) error {
			_ = w.conn.SetReadDeadline(time.Now().Add(po.pongTimeout))
			log.Println("[websockit] pong recv")
			return nil
		})
	} else {
		w.conn.SetPongHandler(func(string) error {
			_ = w.conn.SetReadDeadline(time.Now().Add(po.pongTimeout))
			return nil
		})
	}

	w.pingEnabled = true
}

// PingWithPongTimeout sets a maximum pongTimeout allowed before the Websocket read exits with a read failure
func PingWithPongTimeout(pongTimeout time.Duration) ClientPingOption {
	return func(p *pingOpts) {
		p.pongTimeout = pongTimeout
	}
}

// PingWithPongLog enables client logs when pong messages are received from the server end of the Websocket connection
func PingWithPongLog(logPongs bool) ClientPingOption {
	return func(p *pingOpts) {
		p.logPongs = true
	}
}

// ServerPingLoop infinitely loops, sending new ping messages at a fractional rate of the pongTimeout (default 60s).
// If an error is returned, the client application should discard the current Websocket connection, as any new writes
// will be discarded, and new reads will return the same error. Because this function is blocking, it is recommended to
// wrap it into a goroutine
func (w *WebsocketClient) ServerPingLoop() error {
	if !w.pingEnabled {
		return ErrPingNotEnabled
	}
	pingPeriod := (w.pongTimeout * 9) / 10
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		w.conn.Close()
	}()

	for range ticker.C {
		if err := w.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(5*time.Second)); err != nil {
			return err
		}
	}
	return nil
}
