package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func GetPosts(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"Title": "First Post",
		"Body":  "Hello World!",
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

	r.GET("/posts", GetPosts)

	r.GET("/demo", demoHandler)

	r.GET("/", rootHandler)

	r.POST("/admin/post", postHandler)

	r.POST("/somePost", posting)

	r.GET("/user/:name/*action", userHandler)

	r.GET("/ping", pingHandler)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func rootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Title": "Posts",
		"Article": "This is a post",
		"Text": "Content about your article...",
		"Menu": []string{"Home", "About", "Contact"},
	})
}

func postHandler (c *gin.Context) {

	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("name: %s; message: %s", name, message)
}

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

func userHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func demoHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.tmpl", gin.H{
		"Title": "Users",
	})
}
