package auth

import (
	"net/http"

	"github.com/downdelving/backend/internal/server/api/auth/delete"
	"github.com/downdelving/backend/internal/server/api/auth/login"
	"github.com/downdelving/backend/internal/server/api/auth/logout"
	"github.com/downdelving/backend/internal/server/api/auth/refresh"
	"github.com/downdelving/backend/internal/server/api/auth/register"
	"github.com/downdelving/backend/internal/server/api/auth/update"
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIdGenerator util.AccountIdGenerator) {
	registerHandler := register.NewHandler(accountStorage, passwordHasher, accountIdGenerator)
	mux.HandleFunc("POST /api/auth/delete", delete.Post)
	mux.HandleFunc("POST /api/auth/login", login.Post)
	mux.HandleFunc("GET /api/auth/logout", logout.Get)
	mux.HandleFunc("POST /api/auth/refresh", refresh.Post)
	mux.HandleFunc("POST /api/auth/register", registerHandler.Post)
	mux.HandleFunc("POST /api/auth/update", update.Post)
}
