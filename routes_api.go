package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func apiRoutes(r *mux.Router) {
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/auth/registration", RegistrationMaster).Methods("POST")
	r.HandleFunc("/auth/registration/", RegistrationMaster).Methods("POST")
	r.HandleFunc("/auth/login", AuthMaster).Methods("POST")
	r.HandleFunc("/auth/login/", AuthMaster).Methods("POST")
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}