package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)

var (
	hd  = hashids.NewData()
	err error
	id  int
)

func init() {
	initConfig()
	initDb()

	hd.MinLength = config.Hashids.MinLen
	hd.Salt = config.Hashids.Salt
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
		hash := c.Param("hash")
		id, err = parseHash(hash[1:])
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		profile := getProfile(id)

		if profile.Global {
			c.HTML(http.StatusOK, "global.tmpl.js", gin.H{"server": "127.1:8888"})
		}

	})

	// listen and serve
	r.Run(config.HttpHostPort)
}

func parseHash(hash string) (int, error) {
	if len(hash) < config.Hashids.MinLen {
		return 0, errors.New("illegal user\n")
	}

	id := hashids.NewWithData(hd).Decode(hash)
	if len(id) != 1 {
		return 0, errors.New("illegal user\n")
	}

	return id[0], nil
}
