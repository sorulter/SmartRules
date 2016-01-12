package main

import (
	c "github.com/gocubes/config"
)

var (
	config struct {
		BaseDomain   string   `json:"base_domain"`
		HttpHostPort string   `json:"http_host_port"`
		DefaultList  []string `json:"default_list"`
		DSN          string
		Hashids      struct {
			Salt   string `json:"salt"`
			MinLen int    `json:"min_len"`
		}
	}
)

func initConfig() {
	provider, _ := c.New("etc/config.json", "json")
	provider.Get(&config)
}
