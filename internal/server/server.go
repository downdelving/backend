package server

import (
	"fmt"
	"net/http"
)

func StartServer(port int) {
	mux := http.NewServeMux()
	SetupServer(mux)
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, mux)
}
