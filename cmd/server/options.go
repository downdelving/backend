package main

import (
	"log/slog"
)

type Options struct {
	Listen     bool
	ListenPort int
	LogLevel   slog.Level
}
