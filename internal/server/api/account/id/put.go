package id

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	var req PutRequest
	accountID := r.PathValue("id")
	if accountID == "" {
		http.Error(w, "Missing account id", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Missing request body", http.StatusBadRequest)
		return
	}
	account, err := h.AccountStorage.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Email == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	account.Username = req.Username
	account.Email = req.Email
	if err := h.AccountStorage.UpdateAccount(account); err != nil {
		http.Error(w, "Failed to update account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`{"message": "Successfully updated account!"}`))
}
