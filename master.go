package main

import (
	//"encoding/json"
	"net/http"
	//"strconv"
	"encoding/json"
	"github.com/gorilla/schema"
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

// User's registration request
type MasterRegistrationRequest struct {
	Firstname string   `json:"firstname" schema:"firstname"`
	Lastname  string   `json:"lastname" schema:"lastname"`
	Email 	  string   `json:"email" schema:"email"`
	Phone	  string   `json:"phone" schema:"phone"`
	Password  string   `json:"password" schema:"password"`
}
// Registration response
type RegistrationResponse struct {
	Status 	bool 	`json:"status"`
	Message string 	`json:"message"`
}

// Registration user
func RegistrationMaster(w http.ResponseWriter, r *http.Request) {
	var masterReq MasterRegistrationRequest
	if r.Header.Get("Content-type") == "application/json" {
		json.NewDecoder(r.Body).Decode(&masterReq)
	} else {
		r.ParseForm()
		schemaDec := schema.NewDecoder()
		schemaDec.Decode(&masterReq, r.PostForm)
	}

	if (masterReq.Phone == "") {
		res := RegistrationResponse{
			false,
			"Phone is incorrect",
		}

		json.NewEncoder(w).Encode(res)
	}

	resp := BitrixSearchUser(masterReq.Phone)

	if resp.Total > 0 {
		w.Write([]byte(`{"status": "true"}`))
	} else {
		w.Write([]byte(`{"status": "false"}`))
	}
}



