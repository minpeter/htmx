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

	r.GET("/timedqueue", func(c *gin.Context) {
		RenderTemplates(c, gin.H{})
	})

	timedQueue := NewTimedQueue()

	// var timer *time.Timer
	// var stime time.Time

	r.GET("/timedqueue/sse", func(c *gin.Context) {

		for {

			status := timedQueue.getQueueStatus()

			c.SSEvent("progress", status)
			c.Writer.Flush()

			time.Sleep(time.Millisecond * 500)
		}

		// timer가 몇초가 지났는지 리턴한다
		// for {
		// 	elapsed := time.Since(stime)
		// 	c.SSEvent("progress", fmt.Sprintf("%f", elapsed.Seconds()))
		// 	c.Writer.Flush()
		// 	time.Sleep(time.Millisecond * 500)
		// }

	})

	r.GET("/timedqueue/add", func(c *gin.Context) {

		itme := c.Query("item")

		timedQueue.enqueue(itme)

		// stime = time.Now()
		// time.AfterFunc(10*time.Second, func() {
		// 	fmt.Println("10초가 지났습니다.")
		// })

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
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
