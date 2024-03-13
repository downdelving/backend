package api

import (
	"net/http"

	"github.com/downdelving/backend/internal/server/api/account"
	"github.com/downdelving/backend/internal/server/api/auth"
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIdGenerator util.AccountIdGenerator) {
	auth.SetupMux(mux, accountStorage, passwordHasher, accountIdGenerator)
	account.SetupMux(mux, accountStorage, passwordHasher)
}
