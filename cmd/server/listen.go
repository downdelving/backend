package main

import (
	"flag"
	"log/slog"
	"os"
)

func ParseListenCommand(options *Options) {
	listenCmd := flag.NewFlagSet("listen", flag.ExitOnError)
	listenPortPtr := listenCmd.Int("port", 8080, "port number for server to listen on")
	listenCmd.Parse(os.Args[2:])
	options.Listen = true
	options.ListenPort = *listenPortPtr
	options.LogLevel = slog.LevelInfo
}
