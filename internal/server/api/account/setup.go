package account

import (
	"net/http"

	accountId "github.com/downdelving/backend/internal/server/api/account/id"
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher) {
	accountIdHandler := accountId.NewHandler(accountStorage, passwordHasher)
	mux.HandleFunc("DELETE /api/account/{id}", accountIdHandler.Delete)
}
