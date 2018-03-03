package models

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"strconv"
)

// The person Type (more like an object)
type Master struct {
	ID        int      `json:"id"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Email 	  string   `json:"email"`
}


var masters = []Master{
	Master{ID: 1, Firstname: "Alex", Lastname: "Lebedev", Email: "test@mail.ru"},
	Master{ID: 2, Firstname: "Denis", Lastname: "HedgeHog", Email: "ez@mail.ru"},
}

// Display all from the people var
func GetMasters(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(masters)
}

// Display a single data
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
