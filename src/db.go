package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lessos/lessgo/logger"
)

var (
	db gorm.DB
)

func initDb() {
	if config.DSN == "" {
		fmt.Println("[init db]Must config mysql dsn.")
		os.Exit(0)
	}
	// mysql
	db, err = gorm.Open("mysql", config.DSN)
	if err != nil {
		logger.Print("warn", err.Error())
		time.Sleep(1e9)
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)
	// db.LogMode(true)

}
