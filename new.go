package websockit

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// WebsocketOption specifies any custom options the websocket connection should have
type WebsocketOption func(*Websocket)

// NewWebsocket should be called for any new websocket pair, server or client. This will initialize the websocket
// struct which will be used to accept any websocket options that are set
func NewWebsocket() *Websocket {
	return &Websocket{
		conn:   nil,
		dialer: &websocket.Dialer{},
	}
}

// ServerSocket sets up a new server websocket
func (w *Websocket) ServerSocket(connectUrl string, requestHeader http.Header, opts ...WebsocketOption) error {
	w.setSocketOpts()
	return w.dial(connectUrl, requestHeader)
}

// ClientSocket sets up a new client websocket
func (w *Websocket) ClientSocket(connectUrl string, requestHeader http.Header, opts ...WebsocketOption) error {
	w.setSocketOpts()
	w.dialer.TLSClientConfig = nil
	return w.dial(connectUrl, requestHeader)
}

func (w *Websocket) setSocketOpts(opts ...WebsocketOption) {
	for _, o := range opts {
		o(w)
	}
}

func (w *Websocket) dial(connectUrl string, requestHeader http.Header) error {
	conn, _, err := w.dialer.Dial(connectUrl, requestHeader)
	w.conn = conn
	return err
}

// WithProxy takes a proxy func and runs each new http.Request through this func
func WithProxy(h func(*http.Request) (*url.URL, error)) WebsocketOption {
	return func(w *Websocket) {
		w.dialer.Proxy = h
	}
}

// WithTLSConfig should only be used for websocket servers. If you want to enable encrypted websockets (wss),
// set a TLS certificate chain in the tls.Config
func WithTLSConfig(t *tls.Config) WebsocketOption {
	return func(w *Websocket) {
		w.dialer.TLSClientConfig = t
	}
}

// WithHandshakeTimeout sets a timeout duration for the websocket handshake
func WithHandshakeTimeout(t time.Duration) WebsocketOption {
	return func(w *Websocket) {
		w.dialer.HandshakeTimeout = t
	}
}

// WithReadBufferSize sets the size limit (in bytes) of read buffers in the websocket
func WithReadBufferSize(bufferSize int) WebsocketOption {
	return func(w *Websocket) {
		w.dialer.ReadBufferSize = bufferSize
	}
}

// WithWriteBufferSize sets the size limit (in bytes) of write buffers in the websocket
func WithWriteBufferSize(bufferSize int) WebsocketOption {
	return func(w *Websocket) {
		w.dialer.WriteBufferSize = bufferSize
	}
}

// WithSubprotocols should be used to set the client's preferred subprotocols
func WithSubprotocols(protocols []string) WebsocketOption {
	return func(w *Websocket) {
		w.dialer.Subprotocols = protocols
	}
}
