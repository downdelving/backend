package net

import (
	"net/http"

	rest "github.com/downdelving/backend/internal/net/rest"
	websocket "github.com/downdelving/backend/internal/net/websocket"
)

func SetupServer(mux *http.ServeMux) {
	rest.SetupMux(mux)
	websocket.SetupMux(mux)
}
