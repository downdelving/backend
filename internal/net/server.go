package net

import (
	"net/http"

	rest "github.com/downdelving/backend/internal/net/rest"
	websocket "github.com/downdelving/backend/internal/net/websocket"
)

func StartServer() {
	mux := http.NewServeMux()
	rest.SetupMux(mux)
	websocket.SetupMux(mux)
	http.ListenAndServe(":8080", mux)
}
