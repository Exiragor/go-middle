package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

// Bitrix User
type Bitrixuser struct {
	ID 				string 	`json:"ID"`
	Active 			bool 	`json:"ACTIVE"`
	Email 			string 	`json:"EMAIL"`
	Name 			string 	`json:"NAME"`
	LastName 		string 	`json:"LAST_NAME"`
	PersonalMobile 	string 	`json:"PERSONAL_MOBILE"`
}

//response for BitrixSearchUser
type BitrixSearchUserResponse struct {
	Result 		[]Bitrixuser 	`json:"result"`
	Total 		int 			`json:"total"`
}

//request for add new user
type BitrixAddUserReq struct {
	Email 		string 	`json:"EMAIL"`
	Name 		string 	`json:"NAME"`
	LastName 	string 	`json:"LAST_NAME"`
	Phone 		string 	`json:"PERSONAL_MOBILE"`
	Active 		bool 	`json:"ACTIVE"`
	Extranet 	string 	`json:"EXTRANET"`
	GroupID 	[]int 	`json:"SONET_GROUP_ID"`
}

//response for BitrixAddUser
type BitrixAddUserResponse struct {
	Result int `json:"result"`
}

// Search user in bitrix24
func BitrixSearchUser(phone string) BitrixSearchUserResponse {
	url := Conf.BitrixHook + "user.search.json"

	jsonReq := []byte(`{"PERSONAL_MOBILE": "` + phone + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	response := BitrixSearchUserResponse{}

	json.Unmarshal(body, &response)

	return response
}

// add new use in bitrix24
func BitrixAddUser(master *Master) {
	url := Conf.BitrixHook + "user.add.json"

	requestData := BitrixAddUserReq{
		master.Email,
		master.Firstname,
		master.Lastname,
		master.Phone,
		true,
		"Y",
		[]int{11},
	}
	jsonReq, _ := json.Marshal(requestData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var response BitrixAddUserResponse

	json.Unmarshal(body, &response)

	master.BitrixID = response.Result
}
