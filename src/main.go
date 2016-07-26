package main

import (
	"encoding/json"
	"errors"
	"flag"
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

	prefix string
)

func init() {
	flag.StringVar(&prefix, "prefix", "./", "prefix path")
	flag.Parse()
	initConfig()
	initDb()

	hd.MinLength = config.Hashids.MinLen
	hd.Salt = config.Hashids.Salt
}

func main() {
	defer db.Close()

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
	r.LoadHTMLGlob(prefix + "etc/*")
	r.GET("/v1/*hash", func(c *gin.Context) {
		hash := c.Param("hash")
		id, err = parseHash(hash[1:])
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		profile := getProfile(id)
		server := fmt.Sprintf("%s.%s:%d", profile.NodeName, config.BaseDomain, profile.Port)

		if profile.Global {
			c.HTML(http.StatusOK, "global.tmpl.js", gin.H{"server": server})
		}

		if !profile.Global && profile.UserId != 0 {
			list := func() (data []byte) {
				var rules map[string]int
				json.Unmarshal([]byte(profile.Rules), &rules)
				for k, v := range config.DefaultList {
					rules[k] = v
				}
				data, _ = json.Marshal(rules)

				return
			}()

			c.HTML(http.StatusOK, "pac.tmpl.js", gin.H{
				"server": server,
				"list":   template.HTML(string(list)),
			})
		}

	})

	r.GET("/ping", func(c *gin.Context) {
		var origin = c.Request.Header.Get("Origin")
		for _, cros := range config.CORS {
			if cros == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}

		c.String(http.StatusOK, "pong")
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
