package main

import (
	"bloggin/pkg/logger"
	webserver "bloggin/pkg/web"
)

func main() {
	logger.Msg("INFO","Starting server...")
	webserver.StartServer()
}
