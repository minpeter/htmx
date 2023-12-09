package main

import (
	"net/http"
	"time"

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

	r.GET("/progress", func(c *gin.Context) {
		RenderTemplates(c, gin.H{})
	})

	r.GET("/progress/sse", progressor)

	r.GET("/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.tmpl", nil)
	})

	r.GET("/del", func(c *gin.Context) {
		c.HTML(http.StatusOK, "remove.tmpl", nil)
	})

	r.Run("localhost:8888")
}

func progressor(c *gin.Context) {

	c.SSEvent("progress", "Starting...")
	c.Writer.Flush()

	time.Sleep(1 * time.Second)

	c.SSEvent("progress", "first job completed")
	c.Writer.Flush()

	time.Sleep(1 * time.Second)

	c.SSEvent("progress", "second job completed")
	c.Writer.Flush()

	time.Sleep(1 * time.Second)

	c.SSEvent("progress", "third job completed")
	c.Writer.Flush()

	time.Sleep(1 * time.Second)

	c.SSEvent("progress", "done")

}
