package api

import (
	"net/http"

	auth "github.com/downdelving/backend/internal/server/api/auth"
	"github.com/downdelving/backend/pkg/storage"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.IAccountStorage) {
	auth.SetupMux(mux, accountStorage)
}
