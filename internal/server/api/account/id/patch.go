package id

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	var req PatchRequest
	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Missing request body", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	accountID := r.PathValue("id")
	if accountID == "" {
		http.Error(w, "Missing account id", http.StatusBadRequest)
		return
	}
	account, err := h.AccountStorage.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	if req.Username != "" {
		account.Username = req.Username
	}
	if req.Email != "" {
		account.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := h.PasswordHasher.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		account.Password = hashedPassword
	}
	if err := h.AccountStorage.UpdateAccount(account); err != nil {
		http.Error(w, "Failed to update account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully updated account!"}`))
}
