package storage

import "github.com/downdelving/backend/pkg/model"

type AccountStorage interface {
	CreateAccount(account model.Account) error
	GetAccountByID(id string) (model.Account, error)
	GetAccountByEmail(email string) (model.Account, error)
	GetAccountByUsername(username string) (model.Account, error)
	GetAccountByUsernameOrEmail(usernameOrEmail string) (model.Account, error)
	GetAllAccounts() ([]model.Account, error)
	UpdateAccount(account model.Account) error
	DeleteAccount(id string) error
}
