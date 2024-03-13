package register

import (
	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

type Handler struct {
	AccountStorage     storage.AccountStorage
	PasswordHasher     util.PasswordHasher
	AccountIDGenerator util.AccountIDGenerator
}

func NewHandler(accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIDGenerator util.AccountIDGenerator) *Handler {
	return &Handler{
		AccountStorage:     accountStorage,
		PasswordHasher:     passwordHasher,
		AccountIDGenerator: accountIDGenerator,
	}
}
