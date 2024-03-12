package websocket

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func SetupMux(mux *http.ServeMux) {
	mux.Handle("/ws", websocket.Handler(WebSocketHandler))
}
