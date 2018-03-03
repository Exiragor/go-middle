package main

import (
	"log"
	"net/http"
	"github.com/Exiragor/middle/routes"
	"github.com/Exiragor/middle/database"
)

func main() {
	//init config and routes
	ConfigInit()
	database.DatabaseInit(Conf.Db.Username, Conf.Db.Password, Conf.Db.Tablename)
	router := routes.RoutesInit()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + Conf.App.Port, router))
}

