package webserver

import (
	"bloggin/config/logger"
	"flag"
)

func Flags() {
	server := flag.String("server", "passive", "--server passive")
	flag.Parse()

	if *server != "passive" {
		logger.Msg("WARNING", "Server Mode: "+*server)
	}
}
