package websockit

import (
	"bytes"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

// LocalAddr returns the local network address
func (w *Websocket) LocalAddr() net.Addr {
	return w.conn.LocalAddr()
}

// RemoteAddr returns the remote network address
func (w *Websocket) RemoteAddr() net.Addr {
	return w.conn.RemoteAddr()
}

// CloseNice closes the Websocket connection by sending a CloseNormalClosure control message to the other end of the connection
func (w *Websocket) CloseNice() {
	msg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	w.conn.WriteControl(websocket.CloseMessage, msg, time.Now().Add(3*time.Second))
}

// CloseNow closes the Websocket connection without sending a Close control message
func (w *Websocket) CloseNow() error {
	return w.conn.Close()
}

// ReadLoop infinitely loops, reading the current Websocket for incoming messages and inserting them into `readChannel`.
// If an error is returned, the client application should discard the current Websocket connection, as any new writes
// will be discarded, and new reads will return the same error. Because this function is blocking, it is recommended to
// wrap it into a goroutine
func (w *Websocket) ReadLoop(readChannel chan []byte) error {
	for {
		_, message, err := w.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				return err
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
		readChannel <- message
	}
	return nil
}

// WriteBinaryMessage writes a binary data message to the Websocket
func (w *Websocket) WriteBinaryMessage(data []byte) error {
	return w.conn.WriteMessage(websocket.BinaryMessage, data)
}

// WriteTextMessage writes a UTF-8 encoded interpreted message to the Websocket
func (w *Websocket) WriteTextMessage(data []byte) error {
	return w.conn.WriteMessage(websocket.TextMessage, data)
}
