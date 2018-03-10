package main

import (
	//"encoding/json"
	"net/http"
	//"strconv"
)

// The master Type
type Master struct {
	ID        int      `json:"id"`
	BitrixID  int      `json:"-"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Email 	  string   `json:"email"`
	Phone	  string   `json:"phone"`
	Password  string   `json:"-"`
}

// Display all masters
//func GetMasters(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(masters)
//}
//
//// Display a single master
//func GetMaster(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	for _, item := range masters {
//		id, _ := strconv.Atoi(params["id"])
//
//		if item.ID == id {
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(&Master{})
//}

// Registration user
func RegistrationUser(w http.ResponseWriter, r *http.Request) {
	master := Master{1, 1, "Вася", "Иванов", "l_sf@mail.ru", "+7", "pass123"}
	BitrixSearchUser(master.Email)
}

