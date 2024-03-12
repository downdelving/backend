package net

import (
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	SetupServer(mux)
	http.ListenAndServe(":8080", mux)
}
