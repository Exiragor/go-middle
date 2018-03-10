package main

import (
	"github.com/gorilla/mux"
	//"github.com/justinas/alice"
)

func RoutesInit() *mux.Router {
	r := mux.NewRouter()

	apiRoutes(r)

	return r
}
