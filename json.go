package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	height float64
	wright float64
}

type Person struct {
	Name     string   `json: name`
	Appid    int      `json: appid`
	Info     Info     `json: info`
	Relative []string `json: relative`
}

func main() {
	p1 := `{"appid": 112,
	"name":"natzhu",
	"info": {
	"height": 180.1,
	"weight": 140.3},
	"relative": [
		"father",
		"mother",
		"brother"]
	}`
	p1_binary := []byte(p1)

	var p1_struct map[string]interface{}
	//var p1_struct Person
	err := json.Unmarshal(p1_binary, &p1_struct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1_struct["info"])
}
