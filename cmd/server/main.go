package main

import (
	"fmt"
	"log/slog"

	server "github.com/downdelving/backend/internal/server"
)

func main() {
	options := Options{}
	parseFlags(&options)
	slog.SetLogLoggerLevel(options.LogLevel)
	if options.Listen {
		slog.Info(fmt.Sprintf("Starting server on port %d...", options.ListenPort))
		server.StartServer(options.ListenPort)
	}
}
