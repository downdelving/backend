package main

import (
	"fmt"
	"log/slog"

	"github.com/downdelving/backend/internal/server"
	"github.com/downdelving/backend/internal/storage"
)

func main() {
	options := Options{}
	parseFlags(&options)
	slog.SetLogLoggerLevel(options.LogLevel)
	if options.Listen {
		slog.Info(fmt.Sprintf("Starting server on port %d...", options.ListenPort))
		accountStorage := storage.NewInMemoryAccountStorage()
		server.StartServer(options.ListenPort, accountStorage)
	}
}
