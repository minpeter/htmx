package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func viewRender(c *gin.Context) {
	var view []byte

	templateName := c.Request.URL.Path

	if templateName == "/" {
		templateName = "home"
	}

	if c.GetHeader("Hx-Request") == "true" {
		view, _ = RenderTemplates("htmx", templateName)
	} else {
		view, _ = RenderTemplates("main", templateName)
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", view)
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/components/*")

	view := r.Group("/")
	view.Use(viewRender)
	view.GET("/", func(c *gin.Context) {})
	view.GET("/challenge", func(c *gin.Context) {})

	r.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.tmpl", nil)
	})

	r.GET("/del", func(c *gin.Context) {
		c.HTML(http.StatusOK, "remove.tmpl", nil)
	})

	r.Run("localhost:8080")
}
