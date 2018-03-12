package main

import (
	//"encoding/json"
	"net/http"
	//"strconv"
	"encoding/json"
	"github.com/gorilla/schema"
	"reflect"
	"strconv"
)

//constants


// The master Type
type Master struct {
	ID        int      `json:"id" gorm:"PRIMARY_KEY"`
	BitrixID  int      `json:"-"`
	Firstname string   `json:"firstname" schema:"firstname"`
	Lastname  string   `json:"lastname" schema:"lastname"`
	Email 	  string   `json:"email" schema:"email"`
	Phone	  string   `json:"phone" gorm:"unique; not nul" schema:"phone"`
	Password  string   `json:"-" schema:"password"`
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

// Registration response
type RegistrationResponse struct {
	Status 	bool 	`json:"status"`
	Message string 	`json:"message"`
}

// Registration master
func RegistrationMaster(w http.ResponseWriter, r *http.Request) {
	var master Master
	if r.Header.Get("Content-type") == "application/json" {
		json.NewDecoder(r.Body).Decode(&master)
	} else {
		r.ParseForm()
		schemaDec := schema.NewDecoder()
		schemaDec.Decode(&master, r.PostForm)
	}

	// required fields for registration
	requiredFields := []string{"Phone", "Password"}

	v := reflect.ValueOf(master)
	strIncompleteElems := ""
	for _, elem := range requiredFields {
		value := v.FieldByName(elem).Interface()
		if value == "" {
			strIncompleteElems += elem + " is incorrect; "
		}
	}

	if strIncompleteElems != "" {
		res := RegistrationResponse{
			false,
			"Not all required fields is complete: " + strIncompleteElems,
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	// search master with phone on our db
	var smaster Master
	Db.Where("phone = ?", master.Phone).First(&smaster)

	if smaster.Phone != "" {

		resp := RegistrationResponse{
			false,
			"This phone is already taken",
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	// try to find master in users of bitrix24 on field - phone
	resp := BitrixSearchUser(master.Phone)
	if resp.Total > 0 {
		id, _ := strconv.Atoi(resp.Result[0].ID)
		master.BitrixID = id
	} else {
		BitrixAddUser(&master)
	}

	master.Password, _ = HashPassword(master.Password)

	Db.Create(&master)

	response := RegistrationResponse{
		true,
		"New master is created",
	}

	json.NewEncoder(w).Encode(response)
	return
}



