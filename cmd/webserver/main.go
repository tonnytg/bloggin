package main

import (
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

	// Start Web server GIN
	logger.Msg("INFO", "Starting server...")
	webserver.StartServer()
}
