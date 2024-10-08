// Package handlers handlers package contains gin handlers
package handlers

import (
	"bloggin/internal/entity/post"
	"bloggin/internal/infra/database"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

var (
	Menu = []string{"home", "articles", "about", "contact", "admin"}
)

func RootHandler(c *gin.Context) {

	articles := database.GetAllArticles()

	fmt.Println("valor", len(articles))

	if len(articles) <= 0 {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title":    "Empty",
			"Text":     "Empty",
			"Menu":     Menu,
			"Articles": articles,
		})
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title":    articles[0].Title,
			"Text":     articles[0].Text,
			"Menu":     Menu,
			"Articles": articles,
		})
	}

}

func AdminHandler(c *gin.Context) {
	articles := database.GetAllArticles()

	data := gin.H{
		"Title":    "Admin Panel",
		"Menu":     Menu,
		"Articles": articles,
	}

	if articles == nil {
		data["Title"] = "No Articles"
		data["Text"] = "No articles available."
	}

	c.HTML(http.StatusOK, "admin.tmpl", data)
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

	location := url.URL{Path: "/"}
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
