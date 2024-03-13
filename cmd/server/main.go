package main

import (
	"fmt"
	"log/slog"

	"github.com/downdelving/backend/internal/server"
	"github.com/downdelving/backend/internal/storage/inmemory/accountstorage"
	"github.com/downdelving/backend/internal/util/account/idgenerator"
	"github.com/downdelving/backend/internal/util/account/passwordhasher"
)

func main() {
	options := Options{}
	parseFlags(&options)
	slog.SetLogLoggerLevel(options.LogLevel)
	if options.Listen {
		slog.Info(fmt.Sprintf("Starting server on port %d...", options.ListenPort))
		accountStorage := accountstorage.New()
		passwordHasher := passwordhasher.NewBcrypt()
		accountIdGenerator := &idgenerator.Uuid{}
		server.StartServer(options.ListenPort, accountStorage, passwordHasher, accountIdGenerator)
	}
}
