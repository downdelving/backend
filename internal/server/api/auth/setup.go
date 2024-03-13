package api

import (
	"net/http"

	"github.com/downdelving/backend/internal/server/api/auth/delete"
	"github.com/downdelving/backend/internal/server/api/auth/login"
	"github.com/downdelving/backend/internal/server/api/auth/logout"
	"github.com/downdelving/backend/internal/server/api/auth/refresh"
	"github.com/downdelving/backend/internal/server/api/auth/register"
	"github.com/downdelving/backend/internal/server/api/auth/update"
	"github.com/downdelving/backend/pkg/storage"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.IAccountStorage) {
	registerHandler := register.NewHandler(accountStorage)
	mux.HandleFunc("POST /api/auth/delete", delete.Post)
	mux.HandleFunc("POST /api/auth/login", login.Post)
	mux.HandleFunc("GET /api/auth/logout", logout.Get)
	mux.HandleFunc("POST /api/auth/refresh", refresh.Post)
	mux.HandleFunc("POST /api/auth/register", registerHandler.Post)
	mux.HandleFunc("POST /api/auth/update", update.Post)
}
