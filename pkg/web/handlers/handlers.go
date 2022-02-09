// Package handlers handlers package contains gin handlers
package handlers

import (
	"bloggin/entity/post"
	"bloggin/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

func RootHandler(c *gin.Context) {

	articles := database.GetAllArticles()

	if articles == nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title":    "Empty",
			"Text":     "Empty",
			"Menu":     []string{"Home","Articles", "About", "Contact"},
			"Articles": articles,
		})
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Title":    articles[0].Title,
		"Text":     articles[0].Text,
		"Menu":     []string{"Home","Articles", "About", "Contact"},
		"Articles": articles,
	})
}

func AdminHandler(c *gin.Context) {

	articles := database.GetAllArticles()

	if articles == nil {
		c.HTML(http.StatusOK, "admin.tmpl", gin.H{
			"Title":    "Empty",
			"Text":     "Empty",
			"Menu":     []string{"Home","Articles", "About", "Contact"},
			"Articles": articles,
		})
	}

	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"Title":    articles[0].Title,
		"Text":     articles[0].Text,
		"Menu":     []string{"Home","Articles", "About", "Contact"},
		"Articles": articles,
	})
}

func PostHandler(c *gin.Context) {

	name := c.PostForm("name")
	message := c.PostForm("message")

	article := post.Post{
		Title: name,
		Text:  message,
	}

	database.SavePost(&article)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})

	location := url.URL{Path: "/",}
	c.Redirect(http.StatusFound, location.RequestURI())

	//c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func Posting(c *gin.Context) {
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

func UserHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DemoHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.tmpl", gin.H{
		"Title": "Users",
	})
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

func DeletePosts(c *gin.Context) {

	q := c.Query("id")

	database.DeletePost(q)

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
