package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lessos/lessgo/logger"
)

type Profile struct {
	NodeName string
	Port     int
	UserId   int
	Rules    string
	Global   bool
}

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

func getProfile(id int) (profile Profile) {
	db.Table("ports").Select(
		[]string{"`port`", "node_name", "ports.user_id", "pacs.`global`", "pacs.rules"}).Joins(
		"INNER JOIN pacs ON pacs.user_id = ports.user_id").Where("ports.user_id = ? ", id).Find(&profile)
	fmt.Printf("pac: %v\n", profile)
	return
}
