package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
)

//response for BitrixSearchUser

// Search user in bitrix24
func BitrixSearchUser(email string) []byte {
	url := Conf.BitrixHook + "user.search.json"

	jsonReq := []byte(`{"EMAIL": "` + email + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return body
}
