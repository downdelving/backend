package register

import (
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

type Handler struct {
	AccountStorage     storage.AccountStorage
	PasswordHasher     util.PasswordHasher
	AccountIdGenerator util.AccountIdGenerator
}

func NewHandler(accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIdGenerator util.AccountIdGenerator) *Handler {
	return &Handler{
		AccountStorage:     accountStorage,
		PasswordHasher:     passwordHasher,
		AccountIdGenerator: accountIdGenerator,
	}
}
