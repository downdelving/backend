package server

import (
	"net/http"

	api "github.com/downdelving/backend/internal/server/api"
	websocket "github.com/downdelving/backend/internal/server/websocket"
	"github.com/downdelving/backend/pkg/storage"
)

func SetupServer(mux *http.ServeMux, accountStorage storage.IAccountStorage) {
	api.SetupMux(mux, accountStorage)
	websocket.SetupMux(mux)
}
