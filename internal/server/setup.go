package server

import (
	"net/http"

	rest "github.com/downdelving/backend/internal/server/rest"
	websocket "github.com/downdelving/backend/internal/server/websocket"
)

func SetupServer(mux *http.ServeMux) {
	rest.SetupMux(mux)
	websocket.SetupMux(mux)
}
