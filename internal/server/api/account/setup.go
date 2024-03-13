package account

import (
	"net/http"

	accountID "github.com/downdelving/backend/internal/server/api/account/id"
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIDGenerator util.AccountIDGenerator) {
	accountHandler := NewHandler(accountStorage, passwordHasher, accountIDGenerator)
	accountIDHandler := accountID.NewHandler(accountStorage, passwordHasher)
	mux.HandleFunc("POST /api/account", accountHandler.Post)
	mux.HandleFunc("GET /api/account/{id}", accountIDHandler.Get)
	mux.HandleFunc("PATCH /api/account/{id}", accountIDHandler.Patch)
	mux.HandleFunc("PUT /api/account/{id}", accountIDHandler.Put)
	mux.HandleFunc("DELETE /api/account/{id}", accountIDHandler.Delete)
}
