package main

import (
	"os"
	"encoding/json"
	"fmt"
)

type Configuration struct {
	App			App
	Db			Database
	BitrixHook 	string
}

type App struct {
	Name	string
	Port	string
}

type Database struct {
	Username	string
	Password	string
	Name		string
}

var Conf Configuration

func ConfigInit() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)

	if err != nil {
		fmt.Println("Config error:", err)
	}
}
