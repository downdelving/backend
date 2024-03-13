package id

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
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
	response := GetResponse{
		Account: account,
	}
	marshalledResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledResponse)
}
