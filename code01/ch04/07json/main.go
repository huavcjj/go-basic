package main

import (
	"encoding/json"
	"log"
)

var content = `{
	"species": "pigeon",
	"description": "likes to perch on rocks",
	"dimensions": {
		"height": 24,
		"width": 10
	}
}`

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func main() {
	var data Data
	err := json.Unmarshal([]byte(content), &data)
	// err := json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

}
