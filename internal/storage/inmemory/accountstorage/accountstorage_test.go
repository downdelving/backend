package accountstorage_test

import (
	"errors"
	"testing"

	"github.com/downdelving/backend/internal/storage/inmemory/accountstorage"
	model "github.com/downdelving/backend/pkg/model"
)

// type AccountStorage interface {
// 	CreateAccount(account model.Account) error
// 	GetAccountById(id string) (model.Account, error)
// 	GetAccountByEmail(email string) (model.Account, error)
// 	GetAllAccounts() ([]model.Account, error)
// 	UpdateAccount(account model.Account) error
// 	DeleteAccount(id string) error
// }

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name          string
		accounts      map[string]model.Account
		account       model.Account
		expectedError error
	}{
		{
			name:     "Create Account",
			accounts: map[string]model.Account{},
			account: model.Account{
				Id:       "1",
				Email:    "user@domain.com",
				Username: "user",
				Password: "password",
			},
		},
		{
			name: "Duplicate Account",
			accounts: map[string]model.Account{
				"1": {
					Id:       "1",
					Email:    "user@domain.com",
					Username: "user",
					Password: "password",
				},
			},
			account: model.Account{
				Id:       "1",
				Email:    "user@domain.com",
				Username: "user",
				Password: "password",
			},
			expectedError: errors.New("account already exists"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := accountstorage.WithAccounts(tt.accounts)
			err := storage.CreateAccount(tt.account)
			if err != nil && tt.expectedError == nil {
				t.Errorf("CreateAccount() error = %v, expected no error", err)
			}
			if err == nil && tt.expectedError != nil {
				t.Errorf("CreateAccount() error = nil, expected error %v", tt.expectedError)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("CreateAccount() error = %v, expected error %v", err, tt.expectedError)
			}
		})
	}
}
