package routes

import (
	"bloggin/pkg/web/handlers"
	"github.com/gin-gonic/gin"
)

// Routes defines the routes for the web application
func Routes(r *gin.Engine) {

	// Set folder templates
	r.LoadHTMLGlob("./pkg/web/templates/**")

	// API return getData
	r.GET("/get", handlers.GetData)
	r.GET("/posts", handlers.GetPosts)
	r.GET("/demo", handlers.DemoHandler)
	r.GET("/", handlers.RootHandler)
	r.POST("/admin/post", handlers.PostHandler)
	r.POST("/somePost", handlers.Posting)
	r.GET("/user/:name/*action", handlers.UserHandler)
	r.GET("/ping", handlers.PingHandler)
}