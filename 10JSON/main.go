package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int	`json:"price"`
	Password string `json:"-"`
	Platform string	`json:"platform"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}


func EncodeJson() {
	courses := []course{
		{"ReactJS Bootcamp",299,"ron123","Udemy",[]string{"web-dev","js"}},
		{"MERN Bootcamp",299,"ron123","Udemy",nil},
	}

	finalJson,err := json.MarshalIndent(courses,"","\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)
}

