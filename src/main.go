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

		out, err := exec.Command("ls").Output()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"stderr": string(out)})
		} else {
			res := []string{"foo", "bar"}
			c.HTML(http.StatusOK, "queue.html", res)
		}

	})

	r.Run(":8080")
}
