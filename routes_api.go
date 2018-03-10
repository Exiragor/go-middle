package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func apiRoutes(r *mux.Router) {
	r.HandleFunc("/", YourHandler)
	//r.HandleFunc("/people", models.GetMasters).Methods("GET")
	//r.HandleFunc("/people/{id}", models.GetMaster).Methods("GET")
	r.HandleFunc("/search", RegistrationUser)
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}