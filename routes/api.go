package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/Exiragor/go-middle/models"
)

func apiRoutes(r *mux.Router) {
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/people", models.GetMasters).Methods("GET")
	r.HandleFunc("/people/{id}", models.GetMaster).Methods("GET")
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}