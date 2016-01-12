package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
