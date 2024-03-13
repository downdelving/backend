package auth

import (
	"net/http"

	"github.com/downdelving/backend/internal/server/api/auth/login"
	"github.com/downdelving/backend/internal/server/api/auth/logout"
	"github.com/downdelving/backend/internal/server/api/auth/refresh"
	"github.com/downdelving/backend/internal/server/api/auth/register"
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIDGenerator util.AccountIDGenerator) {
	registerHandler := register.NewHandler(accountStorage, passwordHasher, accountIDGenerator)
	mux.HandleFunc("POST /api/auth/login", login.Post)
	mux.HandleFunc("GET /api/auth/logout", logout.Get)
	mux.HandleFunc("POST /api/auth/refresh", refresh.Post)
	mux.HandleFunc("POST /api/auth/register", registerHandler.Post)
}
