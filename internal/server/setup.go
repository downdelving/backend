package server

import (
	"net/http"

	"github.com/downdelving/backend/internal/server/api"
	"github.com/downdelving/backend/internal/server/websocket"
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func SetupServer(mux *http.ServeMux, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIDGenerator util.AccountIDGenerator) {
	api.SetupMux(mux, accountStorage, passwordHasher, accountIDGenerator)
	websocket.SetupMux(mux)
}
