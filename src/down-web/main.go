package main

import (
	// Web Framework
	"github.com/gin-gonic/gin"
	"net/http"

	// Log
	// "log"
)

const (
	port = ":8080"
)

func main() {
	router := gin.Default()

	router.GET("/down/:action", func(c *gin.Context) {
		message := "Download Task: %s"
		action := c.Param("action")
		c.String(http.StatusOK, message, action)
	})

	router.Run(port)
}
