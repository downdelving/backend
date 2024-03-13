package login

import (
	"encoding/json"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully logged in!"}`))
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Missing request body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.UsernameOrEmail == "" || req.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	account, err := h.AccountStorage.GetAccountByUsernameOrEmail(req.UsernameOrEmail)
	if err != nil {
		http.Error(w, "Failed to get account", http.StatusUnauthorized)
		return
	}
	hashedPassword := account.Password
	match, err := h.PasswordHasher.ComparePassword(hashedPassword, req.Password)
	if match == false || err != nil {
		http.Error(w, "Invalid username/email or password", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Successfully logged in!"}`))
}
