package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string) // string key -> string value map
	languages["js"] = "Javascript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	languages["go"] = "Golang"

	fmt.Println("map of languages",languages)
	fmt.Println("js :",languages["js"])

	// to delete some value
	delete(languages,"rb")
	fmt.Println("map of languages after deleting rb",languages)

	// looping through map (kind of for each loop-js)
	for key, value := range languages {
		// fmt.Println(key, ":", value)
		fmt.Printf("For key %v : Value is %v\n", key, value)
	}
	//if want to ignore key or value can use _
	for _, value := range languages {
		fmt.Println(value)
	}
}