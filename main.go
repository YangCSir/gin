package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./template/*")

	r.GET("/test", func(c *gin.Context) {
		c.HTML(200, "a.html", gin.H{
			"title": "yanglaosan",
		})
	})

	r.Run()
}
