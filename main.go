package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/components/*")

	view := r.Group("/")
	view.GET("/", func(c *gin.Context) {
		RenderTemplates(c, gin.H{})
	})

	view.GET("/challenge", func(c *gin.Context) {
		RenderTemplates(c, gin.H{
			"Text": "Hello, World!",
		})
	})

	r.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.tmpl", nil)
	})

	r.GET("/del", func(c *gin.Context) {
		c.HTML(http.StatusOK, "remove.tmpl", nil)
	})

	r.Run("localhost:8080")
}
