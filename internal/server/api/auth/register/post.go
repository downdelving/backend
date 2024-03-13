package register

import (
	"encoding/json"
	"net/http"

	"github.com/downdelving/backend/internal/util"
	"github.com/downdelving/backend/pkg/model"
	"github.com/downdelving/backend/pkg/storage"
)

type Handler struct {
	AccountStorage storage.IAccountStorage
}

func NewHandler(accountStorage storage.IAccountStorage) *Handler {
	return &Handler{
		AccountStorage: accountStorage,
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Missing request body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	account := model.Account{
		ID:       util.GenerateID(),
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := h.AccountStorage.CreateAccount(account); err != nil {
		http.Error(w, "Failed to register account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Successfully registered!"}`))
}
