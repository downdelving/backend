package login

import (
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

type Handler struct {
	AccountStorage storage.AccountStorage
	PasswordHasher util.PasswordHasher
}

func NewHandler(accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher) *Handler {
	return &Handler{
		AccountStorage: accountStorage,
		PasswordHasher: passwordHasher,
	}
}
