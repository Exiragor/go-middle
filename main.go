package main

import (
	"log"
	"net/http"
	"github.com/Exiragor/middle/routes"
)

func main() {
	//init config and routes
	ConfigInit()
	router := routes.RoutesInit()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + Conf.App.Port, router))
}

