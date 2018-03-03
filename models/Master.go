package models

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"strconv"
)

// The master Type
type Master struct {
	ID        int      `json:"id"`
	BitrixID  int      `json:"-"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Email 	  string   `json:"email"`
	Password  string   `json:"-"`
}


var masters = []Master{
	Master{ID: 1, Firstname: "Alex", Lastname: "Lebedev", Email: "test@mail.ru", Password: "test"},
	Master{ID: 2, Firstname: "Denis", Lastname: "HedgeHog", Email: "ez@mail.ru", Password: "test123"},
}

// Display all masters
func GetMasters(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(masters)
}

// Display a single master
func GetMaster(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range masters {
		id, _ := strconv.Atoi(params["id"])

		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Master{})
}
