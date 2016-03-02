package main

import (
	c "github.com/gocubes/config"
	"log"
	"syscall"
)

var (
	config struct {
		BaseDomain   string         `json:"base_domain"`
		HttpHostPort string         `json:"http_host_port"`
		DefaultList  map[string]int `json:"default_list"`
		DSN          string
		Hashids      struct {
			Salt   string `json:"salt"`
			MinLen int    `json:"min_len"`
		}
		CORS []string `json:"cros"`
	}
)

func initConfig() {
	provider, er := c.New(prefix+"etc/config.json", "json")
	if er != nil {
		log.Fatalf("read config file error:%s\n", er.Error())
	}
	provider.Get(&config)
	provider.ReloadOn(&config, syscall.SIGUSR1)
}
