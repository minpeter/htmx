package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/components/*")

	r.GET("/", func(c *gin.Context) {
		var view []byte

		// check Hx-Request header
		if c.GetHeader("Hx-Request") == "true" {
			view, _ = RenderTemplates("htmx", "home")
		} else {
			view, _ = RenderTemplates("main", "home")
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", view)
	})

	r.GET("/challenge", func(c *gin.Context) {

		var view []byte

		// check Hx-Request header
		if c.GetHeader("Hx-Request") == "true" {
			view, _ = RenderTemplates("htmx", "challenge")
		} else {
			view, _ = RenderTemplates("main", "challenge")
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", view)
	})

	// GET /new 라우터를 만들고 templates/pages/create.tmpl를 렌더링합니다.
	r.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.tmpl", nil)
	})

	r.GET("/del", func(c *gin.Context) {
		c.HTML(http.StatusOK, "remove.tmpl", nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
