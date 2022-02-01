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

	// Start Web server GIN
	logger.Msg("INFO", "Starting server...")
	webserver.StartServer()
}
