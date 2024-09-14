package main

import (
	"bloggin/cmd"
	"bloggin/config/logger"
	"bloggin/internal/infra/database"
	"bloggin/internal/infra/web"
	"os"
)

func init() {
	// Database init
	database.InitDatabase()
}

func main() {
	// Set log level from environment variable, default to INFO
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}
	os.Setenv("LOG_LEVEL", logLevel)

	// Check cmd
	cmd.Flags()

	// Start Web server GIN
	logger.Msg("INFO", "Starting server...")
	webserver.StartServer()
}
