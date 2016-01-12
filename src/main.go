package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var (
	err error
	id  int
)

func init() {
	initConfig()
	initDb()
}

func main() {

	go run()

	// Main loop
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
forever:

	for {
		select {
		case s := <-sig:
			fmt.Printf("\nSignal (%d) received, stopping\n", s)
			break forever
		}
	}
}

func run() {
	r := gin.Default()
	r.LoadHTMLGlob("etc/*")
	r.GET("/v1/*hash", func(c *gin.Context) {

	})

	// listen and serve
	r.Run(config.HttpHostPort)
}
