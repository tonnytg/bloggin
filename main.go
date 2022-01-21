package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB string `form:"field_b"`
}

func GetData(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func main() {

	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Set folder templates
	r.LoadHTMLGlob("templates/**")

	// API return getData
	r.GET("/get", GetData)

	r.GET("/demo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo.tmpl", gin.H{
			"Title": "Users",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title": "Posts",
			"Article": "This is a post",
			"Text": "Content about your article...",
			"Menu": []string{"Home", "About", "Contact"},
		})
	})

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
