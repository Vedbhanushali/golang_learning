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

func EncodeJson() []byte {
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
	// fmt.Println(string(finalJson))
	return finalJson
}

func DecodeJson() {
	//consume json
	jsonDataFromWeb := EncodeJson()
	checkValid := json.Valid(jsonDataFromWeb)
	if !checkValid {
		fmt.Println("Json is invalid")
		return
	}
	var mycourses []course 
	json.Unmarshal(jsonDataFromWeb, &mycourses)
	fmt.Printf("%#v\n%#v",mycourses,mycourses[0].Name)

	var dataMap []map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &dataMap) //as decoding array key is index value
	// fmt.Printf("%#v\n",dataMap)
	for k,v := range dataMap {
		fmt.Printf("Key is %v and value is %v\n",k,v)
	}
}

func main() {
	fmt.Println("Encoding Json")
	// EncodeJson()
	DecodeJson()
}