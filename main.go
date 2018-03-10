package main

import (
	"log"
	"net/http"
)

func main() {
	//init config, database and routes
	ConfigInit()
	DatabaseInit(Conf.Db.Username, Conf.Db.Password, Conf.Db.Name)
	router := RoutesInit()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + Conf.App.Port, router))
}

