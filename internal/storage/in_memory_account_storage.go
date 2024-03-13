package storage

import (
	"errors"
	"sync"

	model "github.com/downdelving/backend/pkg/model"
)

// InMemoryAccountStorage is an in-memory storage for account data.
type InMemoryAccountStorage struct {
	mu       sync.RWMutex // Guards users
	accounts map[string]model.Account
}

// NewInMemoryAccountStorage initializes a new InMemoryAccountStorage.
func NewInMemoryAccountStorage() *InMemoryAccountStorage {
	return &InMemoryAccountStorage{
		accounts: make(map[string]model.Account),
	}
}

// CreateAccount stores a new account.
func (s *InMemoryAccountStorage) CreateAccount(account model.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[account.ID]; exists {
		return errors.New("account already exists")
	}

	s.accounts[account.ID] = account
	return nil
}

// GetAccountByID retrieves an account by their ID.
func (s *InMemoryAccountStorage) GetAccountByID(id string) (model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	account, exists := s.accounts[id]
	if !exists {
		return model.Account{}, errors.New("account not found")
	}

	return account, nil
}

// GetAccountByEmail retrieves an account by their email.
func (s *InMemoryAccountStorage) GetAccountByEmail(email string) (model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, account := range s.accounts {
		if account.Email == email {
			return account, nil
		}
	}

	return model.Account{}, errors.New("account not found")
}

// UpdateAccount updates an existing account.
func (s *InMemoryAccountStorage) UpdateAccount(account model.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[account.ID]; !exists {
		return errors.New("account not found")
	}

	s.accounts[account.ID] = account
	return nil
}

// DeleteAccount removes an account from storage.
func (s *InMemoryAccountStorage) DeleteAccount(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[id]; !exists {
		return errors.New("account not found")
	}

	delete(s.accounts, id)
	return nil
}
