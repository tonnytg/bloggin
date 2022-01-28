package main

import (
	"bloggin/pkg/cmd"
	"bloggin/pkg/logger"
	webserver "bloggin/pkg/web"
)

func main() {

	// Check cmd
	cmd.Flags()

	logger.Msg("INFO","Starting server...")

	// Start Web server GIN
	webserver.StartServer()
}
