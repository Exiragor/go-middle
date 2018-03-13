package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/schema"
	"reflect"
	"strconv"
)

//constants

// The Master Type
type Master struct {
	ID        int      `json:"id" gorm:"PRIMARY_KEY"`
	Active    bool     `json:"-"`
	BitrixID  int      `json:"-"`
	Firstname string   `json:"firstname" schema:"firstname"`
	Lastname  string   `json:"lastname" schema:"lastname"`
	Email 	  string   `json:"email" gorm:"unique; not nul" schema:"email"`
	Phone	  string   `json:"phone" gorm:"unique; not nul" schema:"phone"`
	Password  string   `json:"-" schema:"password"`
}

// Registration response
type StatusResponse struct {
	Status 			bool 	`json:"status"`
	BitrixInvite 	bool	`json:"bitrix_invite,omitempty"`
	Message 		string 	`json:"message"`
}

// Registration master
func RegistrationMaster(w http.ResponseWriter, r *http.Request) {
	//parse request
	var master Master
	if r.Header.Get("Content-type") == "application/json" {
		json.NewDecoder(r.Body).Decode(&master)
	} else {
		r.ParseForm()
		schemaDec := schema.NewDecoder()
		schemaDec.Decode(&master, r.PostForm)
	}

	// required fields for registration
	requiredFields := []string{"Phone", "Password", "Email"}

	// validate
	v := reflect.ValueOf(master)
	strIncompleteElems := ""
	for _, elem := range requiredFields {
		value := v.FieldByName(elem).Interface()
		if value == "" {
			strIncompleteElems += elem + " is incorrect; "
		}
	}

	if strIncompleteElems != "" {
		res := StatusResponse{
			Status: false,
			Message: "Not all required fields is complete: " + strIncompleteElems,
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	// search master with phone in our db
	var smaster Master
	Db.Where("phone = ?", master.Phone).First(&smaster)

	if smaster.Phone != "" {
		resp := StatusResponse{
			Status: false,
			Message: "This phone is already taken",
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	// try to find master in users of bitrix24 on field - phone
	resp := BitrixSearchUser(master.Phone)
	bitrixInvite := false

	if resp.Total > 0 {
		id, _ := strconv.Atoi(resp.Result[0].ID)
		master.BitrixID = id
		master.Active = true
	} else {
		BitrixAddUser(&master)
		bitrixInvite = true
		master.Active = false
	}

	master.Password, _ = HashPassword(master.Password)
	Db.Create(&master)

	response := StatusResponse{
		true,
		bitrixInvite,
		"New master was created",
	}

	json.NewEncoder(w).Encode(response)
	return
}



