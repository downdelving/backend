package main

// Import the package that contains the code we want to run from the actual package path
import (
	net "github.com/downdelving/backend/internal/net"
)

func main() {
	net.StartServer()
}
