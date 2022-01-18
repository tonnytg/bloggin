package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func posting (c *gin.Context) {
	var msg string
	if c.Request.Method == http.MethodPost {
		msg = "Post"
	} else {
		// never enter here because filter only POST
		msg = "Get"
	}
	c.JSON(200, gin.H{
		"method": msg,
	})
}

func main() {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	//
	r.SetHTMLTemplate(html)

	r.POST("/somePost", posting)

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
