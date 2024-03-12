package main

import (
	"fmt"
	"os"
)

func parseFlags(options *Options) {
	if len(os.Args) < 2 {
		fmt.Println("expected subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "listen":
		ParseListenCommand(options)
	default:
		fmt.Println("expected 'listen' subcommand")
		os.Exit(1)
	}
}
