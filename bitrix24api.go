package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
)

func BitrixSearchUser(email string) {
	url := Conf.BitrixHook + "user.search.json/"

	fmt.Println(email)

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
}
