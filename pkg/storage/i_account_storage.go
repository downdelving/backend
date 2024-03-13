package storage

import "github.com/downdelving/backend/pkg/model"

type IAccountStorage interface {
	CreateAccount(account model.Account) error
	GetAccountByID(id string) (model.Account, error)
	GetAccountByEmail(email string) (model.Account, error)
	UpdateAccount(account model.Account) error
	DeleteAccount(id string) error
}
