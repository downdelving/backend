package server

import (
	"fmt"
	"net/http"

	"github.com/downdelving/backend/pkg/storage"
)

func StartServer(port int, accountStorage storage.IAccountStorage) {
	mux := http.NewServeMux()
	SetupServer(mux, accountStorage)
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, mux)
}
