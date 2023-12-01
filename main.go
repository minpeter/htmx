package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/challenge", func(c *gin.Context) {
		view, err := RenderTemplates("main", "challenge")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", view)
	})

	// GET /new 라우터를 만들고 templates/pages/create.tmpl를 렌더링합니다.
	r.GET("/new", func(c *gin.Context) {

		view, err := RenderTemplates("main", "create")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", view)
	})

	r.GET("/del", func(c *gin.Context) {

		view, _ := RenderTemplates("main", "remove")

		c.Data(http.StatusOK, "text/html; charset=utf-8", view)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
