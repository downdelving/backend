package server

import (
	"fmt"
	"net/http"

	"github.com/downdelving/backend/pkg/storage"
	"github.com/downdelving/backend/pkg/util"
)

func StartServer(port int, accountStorage storage.AccountStorage, passwordHasher util.PasswordHasher, accountIdGenerator util.AccountIdGenerator) {
	mux := http.NewServeMux()
	SetupServer(mux, accountStorage, passwordHasher, accountIdGenerator)
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, mux)
}
