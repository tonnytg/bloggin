package main

import (
	"bloggin/pkg/cmd"
	"bloggin/pkg/database"
	"bloggin/pkg/logger"
	webserver "bloggin/pkg/web"
)

func init() {
	// Database init
	database.InitDatabase()
}

func main() {

	// Check cmd
	cmd.Flags()

	logger.Msg("INFO", "Starting server...")

	// Start Web server GIN
	webserver.StartServer()
}
