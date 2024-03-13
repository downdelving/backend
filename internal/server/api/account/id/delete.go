package id

import (
	"net/http"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	accountID := r.PathValue("id")
	if accountID == "" {
		http.Error(w, "Missing account id", http.StatusBadRequest)
		return
	}
	_, err := h.AccountStorage.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	if err := h.AccountStorage.DeleteAccount(accountID); err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`{"message": "Successfully deleted account!"}`))
}
