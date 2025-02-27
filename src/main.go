package main

import (
	"net/http"
	"os/exec"

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
			"title": "Kube",
			"ipv4":  "1.1.1.1",
		})
	})

	r.GET("/run", func(c *gin.Context) {

		out, err := exec.Command("ls /app").Output()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"stderr": string(out)})
		} else {
			c.JSON(http.StatusOK, gin.H{"stdout": string(out)})
		}

	})

	r.Run(":8080")
}
