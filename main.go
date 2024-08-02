package main

import (
	"bloggin/pkg/cmd"
	"bloggin/pkg/database"
	"bloggin/pkg/logger"
	webserver "bloggin/pkg/web"
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
