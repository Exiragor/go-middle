package main

import (
	"log"
	"net/http"
	"github.com/Exiragor/go-middle/routes"
	"github.com/Exiragor/go-middle/models"
)

func main() {
	models.SetMasters()
	router := routes.RoutesInit()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":5000", router))
}
