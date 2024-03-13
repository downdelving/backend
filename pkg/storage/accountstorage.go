package storage

import "github.com/downdelving/backend/pkg/model"

type AccountStorage interface {
	CreateAccount(account model.Account) error
	GetAccountById(id string) (model.Account, error)
	GetAccountByEmail(email string) (model.Account, error)
	GetAllAccounts() ([]model.Account, error)
	UpdateAccount(account model.Account) error
	DeleteAccount(id string) error
}
