package webserver

import (
	"bloggin/pkg/web/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// StartServer load gin server with routes, templates and handlers
func StartServer() {

	// Force logger's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Charge templates folder
	r.StaticFS("/public", http.Dir("./pkg/web/public"))
	r.LoadHTMLGlob("./pkg/web/templates/**")

	// Load Routes with same gin process
	routes.Routes(r)

	// Start the server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
