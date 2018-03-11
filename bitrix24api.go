package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

// Bitrix User
type Bitrixuser struct {
	ID 				int 	`json:"ID"`
	Active 			string 	`json:"ACTIVE"`
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
