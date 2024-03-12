package rest

import "net/http"

func handleResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, world!"}`))
}

func SetupMux(mux *http.ServeMux) {
	mux.HandleFunc("/api/resource", handleResource)
}
