package id

import (
	"net/http"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	accountId := r.PathValue("id")
	if accountId == "" {
		http.Error(w, "Missing account id", http.StatusBadRequest)
		return
	}
	if err := h.AccountStorage.DeleteAccount(accountId); err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Successfully deleted account!"}`))
}
