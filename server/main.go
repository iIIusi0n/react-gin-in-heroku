package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func RedirectHttpToHttps(c *gin.Context) {
	if strings.Contains(c.Request.URL.String(), "http://") {
		c.Redirect(http.StatusMovedPermanently, strings.Replace(c.Request.URL.String(), "http://", "https://", -1))
	}

	c.Next()
}

func main() {
	r := gin.Default()
	r.Use(RedirectHttpToHttps)

	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
