package api

import (
	"net/http"

	"github.com/downdelving/backend/internal/server/api/account"
	"github.com/downdelving/backend/internal/server/api/auth"
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func SetupMux(mux *http.ServeMux, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIDGenerator util.AccountIDGenerator) {
	account.SetupMux(mux, accountStorage, passwordHasher, accountIDGenerator)
	auth.SetupMux(mux, accountStorage, passwordHasher, accountIDGenerator)
}
