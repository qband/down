package main

import (
	// Web Framework
	"github.com/gin-gonic/gin"
	"net/http"

	// Command Runner
	"github.com/codeskyblue/go-sh"

	// Log
	"log"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = sh.Command

func main() {
	router := gin.Default()

	router.GET("/down/:action", func(c *gin.Context) {
		message := "Download Task: %s: %s"
		action := c.Param("action")

		out, err := sh.Command("echo", "hello\tworld").Command("cut", "-f2").Output()
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, message, action, out)
	})

	router.Run(":8080")
}

// transmission-daemon
// transmission-remote --add "magnet/torrent_file"
// amulecmd
// amuled
