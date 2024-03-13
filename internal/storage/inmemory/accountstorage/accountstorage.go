package accountstorage

import (
	"errors"
	"sync"

	model "github.com/downdelving/backend/pkg/model"
)

type AccountStorage struct {
	mu       sync.RWMutex // Guards users
	accounts map[string]model.Account
}

func New() *AccountStorage {
	return &AccountStorage{
		accounts: make(map[string]model.Account),
	}
}

func WithAccounts(accounts map[string]model.Account) *AccountStorage {
	return &AccountStorage{
		accounts: accounts,
	}
}

func (s *AccountStorage) CreateAccount(account model.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[account.Id]; exists {
		return errors.New("account already exists")
	}

	s.accounts[account.Id] = account
	return nil
}

func (s *AccountStorage) GetAccountById(id string) (model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	account, exists := s.accounts[id]
	if !exists {
		return model.Account{}, errors.New("account not found")
	}

	return account, nil
}

func (s *AccountStorage) GetAccountByEmail(email string) (model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, account := range s.accounts {
		if account.Email == email {
			return account, nil
		}
	}

	return model.Account{}, errors.New("account not found")
}

func (s *AccountStorage) GetAccountByUsername(username string) (model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, account := range s.accounts {
		if account.Username == username {
			return account, nil
		}
	}

	return model.Account{}, errors.New("account not found")
}

func (s *AccountStorage) GetAccountByUsernameOrEmail(usernameOrEmail string) (model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, account := range s.accounts {
		if account.Username == usernameOrEmail || account.Email == usernameOrEmail {
			return account, nil
		}
	}

	return model.Account{}, errors.New("account not found")
}

func (s *AccountStorage) GetAllAccounts() ([]model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	accounts := make([]model.Account, 0, len(s.accounts))
	for _, account := range s.accounts {
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *AccountStorage) UpdateAccount(account model.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[account.Id]; !exists {
		return errors.New("account not found")
	}

	s.accounts[account.Id] = account
	return nil
}

func (s *AccountStorage) DeleteAccount(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[id]; !exists {
		return errors.New("account not found")
	}

	delete(s.accounts, id)
	return nil
}
