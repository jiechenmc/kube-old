package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Connected": true})
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/clicked", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Connected": true})
	})
	fmt.Println("RUNNING")

	r.Run(":8080")
}
