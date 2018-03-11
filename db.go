package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var Db *gorm.DB

func DatabaseInit(login string, pass string, table string) {
	str := login + ":" + pass + "@/" + table
	dbconn, err := gorm.Open("mysql", str)

	if err != nil {
		fmt.Println("Database error: connection refused")
	}

	Db = dbconn

	autoMigrate()
}

func autoMigrate() {
	Db.AutoMigrate(&Master{})
}