package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/Exiragor/middle/models"
)

var Tool *gorm.DB

func DatabaseInit(login string, pass string, table string) {
	str := login + ":" + pass + "@/" + table
	dbconn, err := gorm.Open("mysql", str)

	if err != nil {
		fmt.Println("Database error: connection refused")
	}

	Tool = dbconn

	autoMigrate()
}

func autoMigrate() {
	Tool.AutoMigrate(&models.Master{})
}