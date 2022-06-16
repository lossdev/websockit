package websockit

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// WebsocketClientOption specifies any custom options a client websocket connection should have
type WebsocketClientOption func(*Websocket)

// WebsocketServerOption specifies any custom options a server websocket connection should have
type WebsocketServerOption func(*Websocket)

// NewWebsocket should be called for any new websocket pair, server or client. This will initialize the websocket
// struct values which will be used to accept any websocket options that are set
func NewWebsocket() *Websocket {
	return &Websocket{
		conn:     nil,
		dialer:   &websocket.Dialer{},
		upgrader: &websocket.Upgrader{},
	}
}

// ServerSocket sets up a new server websocket end
func (w *Websocket) ServerSocket(wr http.ResponseWriter, req *http.Request, headers http.Header, opts ...WebsocketServerOption) (*WebsocketServer, error) {
	for _, o := range opts {
		o(w)
	}
	conn, err := w.upgrader.Upgrade(wr, req, headers)
	if err != nil {
		return nil, err
	}
	w.conn = conn
	return &WebsocketServer{w}, nil
}

// ServerWithHandshakeTimeout sets a timeout duration for the websocket handshake
func ServerWithHandshakeTimeout(t time.Duration) WebsocketServerOption {
	return func(w *Websocket) {
		w.upgrader.HandshakeTimeout = t
	}
}

// ServerWithReadBufferSize sets the size limit (in bytes) of read buffers in the websocket
func ServerWithReadBufferSize(bufferSize int) WebsocketServerOption {
	return func(w *Websocket) {
		w.upgrader.ReadBufferSize = bufferSize
	}
}

// ServerWithWriteBufferSize sets the size limit (in bytes) of write buffers in the websocket
func ServerWithWriteBufferSize(bufferSize int) WebsocketServerOption {
	return func(w *Websocket) {
		w.upgrader.WriteBufferSize = bufferSize
	}
}

// ServerWithSubprotocols should be used to set the server's preferred subprotocols
func ServerWithSubprotocols(protocols []string) WebsocketServerOption {
	return func(w *Websocket) {
		w.upgrader.Subprotocols = protocols
	}
}

// ServerWithErrorFunc can use a custom HTTP error function - otherwise, http.Error will print errors
func ServerWithErrorFunc(errorFunc func(wr http.ResponseWriter, req *http.Request, status int, reason error)) WebsocketServerOption {
	return func(w *Websocket) {
		w.upgrader.Error = errorFunc
	}
}

// ServerWithCheckOriginFunc can override the default CheckOrigin deny policy given the http.Request made to the server
func ServerWithCheckOriginFunc(originFunc func(r *http.Request) bool) WebsocketServerOption {
	return func(w *Websocket) {
		w.upgrader.CheckOrigin = originFunc
	}
}

// ClientSocket sets up a new client websocket end
func (w *Websocket) ClientSocket(connectUrl string, headers http.Header, opts ...WebsocketClientOption) (*WebsocketClient, error) {
	for _, o := range opts {
		o(w)
	}
	conn, _, err := w.dialer.Dial(connectUrl, headers)
	if err != nil {
		return nil, err
	}
	w.conn = conn
	return &WebsocketClient{w, false, 60 * time.Second}, nil
}

// ClientWithProxy takes a proxy func and runs each new http.Request through this func
func ClientWithProxy(proxyFunc func(*http.Request) (*url.URL, error)) WebsocketClientOption {
	return func(w *Websocket) {
		w.dialer.Proxy = proxyFunc
	}
}

// ClientWithTLSConfig sets the TLS config for the websocket session
func ClientWithTLSConfig(t *tls.Config) WebsocketClientOption {
	return func(w *Websocket) {
		w.dialer.TLSClientConfig = t
	}
}

// ClientWithHandshakeTimeout sets a timeout duration for the websocket handshake
func ClientWithHandshakeTimeout(t time.Duration) WebsocketClientOption {
	return func(w *Websocket) {
		w.dialer.HandshakeTimeout = t
	}
}

// ClientWithReadBufferSize sets the size limit (in bytes) of read buffers in the websocket
func ClientWithReadBufferSize(bufferSize int) WebsocketClientOption {
	return func(w *Websocket) {
		w.dialer.ReadBufferSize = bufferSize
	}
}

// ClientWithWriteBufferSize sets the size limit (in bytes) of write buffers in the websocket
func ClientWithWriteBufferSize(bufferSize int) WebsocketClientOption {
	return func(w *Websocket) {
		w.dialer.WriteBufferSize = bufferSize
	}
}

// ClientWithSubprotocols should be used to set the client's preferred subprotocols
func ClientWithSubprotocols(protocols []string) WebsocketClientOption {
	return func(w *Websocket) {
		w.dialer.Subprotocols = protocols
	}
}
