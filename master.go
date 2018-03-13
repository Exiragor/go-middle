package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/schema"
	"strconv"
	"fmt"
	"time"
)

//constants

// The Master Type
type Master struct {
	ID		  	int 			`json:"id"`
	CreatedAt 	*time.Time 		`json:"created_at"`
	UpdatedAt 	*time.Time 		`json:"updated_at"`
	Active    	bool     		`json:"-"`
	BitrixID  	int      		`json:"-"`
	Firstname 	string   		`json:"firstname" schema:"firstname"`
	Lastname  	string   		`json:"lastname" schema:"lastname"`
	Email 	  	string   		`json:"email" gorm:"unique; not nul" schema:"email"`
	Phone	  	string   		`json:"phone" gorm:"unique; not nul" schema:"phone"`
	Password  	string   		`json:"-" gorm:"not null" schema:"password"`
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
	master = parseFields(Master{}, r)

	// validate
	strIncompleteElems := MasterRegistrationValidate(master)

	if strIncompleteElems != "" {
		res := StatusResponse{
			Status: false,
			Message: "Not all required fields is complete: " + strIncompleteElems,
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	// search master with phone in our db
	var smasterPhone Master
	var smasterEmail Master
	Db.Where("phone = ?", master.Phone).First(&smasterPhone)
	Db.Where("email = ?", master.Email).First(&smasterEmail)

	if smasterPhone.Phone != "" {
		resp := StatusResponse{
			Status: false,
			Message: "This phone is already taken",
		}

		json.NewEncoder(w).Encode(resp)
		return
	}
	if smasterEmail.Email != "" {
		resp := StatusResponse{
			Status: false,
			Message: "This email is already taken",
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

// Auth response
type AuthReponse struct {
	ID int
	AccessToken string
	ResponseToken string
}

// Auth master
func AuthMaster(w http.ResponseWriter, r *http.Request) {
	// parse request
	var master Master
	master = parseFields(Master{}, r)

	// validate
	strIncompleteElems := MasterAuthValidate(master)

	if strIncompleteElems != "" {
		res := StatusResponse{
			Status: false,
			Message: "Not all required fields is complete: " + strIncompleteElems,
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	// Try to auth Master
	var smaster Master
	Db.Where("phone = ?", master.Phone).First(&smaster)
	fmt.Println(smaster)
	if smaster.Password == "" {
		res := StatusResponse{
			Status: false,
			Message: "Phone is not found",
		}

		json.NewEncoder(w).Encode(res)
		return
	} else {
		if !CheckPasswordHash(master.Password, smaster.Password) {
			res := StatusResponse{
				Status: false,
				Message: "Password is incorrect",
			}

			json.NewEncoder(w).Encode(res)
			return
		}
	}

	json.NewEncoder(w).Encode(smaster)
	return
}

// parse fields from request for master
func parseFields(master Master, r *http.Request) Master {
	if r.Header.Get("Content-type") == "application/json" {
		json.NewDecoder(r.Body).Decode(&master)
	} else {
		r.ParseForm()
		schemaDec := schema.NewDecoder()
		schemaDec.Decode(&master, r.PostForm)
	}

	return master
}
