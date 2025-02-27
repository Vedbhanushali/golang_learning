package main

import (
	"encoding/json"
	"fmt"
)

type course struct { //private not exporting it
	Name     string `json:"course_name"` //tagging
	Price    int   	//no tagging
	Platform string //no tagging
	Password string `json:"-"` //do parse/show/return this field 
	Tags     []string `json:"tags,omitempty"` //if tags is nil, then don't show it
}

func EncodeJson() {
	courses := []course{
		{"Go", 12, "youtube", "password", []string{"web-dev", "programming"}},
		{"Python", 10, "youtube", "password", []string{"web-dev", "programming", "data-science"}},
		{"Java", 15, "youtube", "password", []string{"web-dev", "programming", "OOP"}},
		{"C++", 20, "youtube", "password", nil},
	}

	//package this data as json
	finalJson, err := json.MarshalIndent(courses,"","\t") // \t is for tab
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

func main() {
	fmt.Println("Encoding Json")
	EncodeJson()
}