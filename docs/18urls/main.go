package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://go.dev:8001/login?type=2&name=ved&password=pass"

func main() {
	fmt.Println("URLs in Go")
	result , _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Query())
	fmt.Println(result.Port())

	queryParams := result.Query()
	fmt.Printf("Type of queryParams %T\n", queryParams)
	fmt.Println(queryParams["name"], queryParams["password"])
	for _ , val := range queryParams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "go.dev",
		Path: "login",
		RawQuery: "type=2&name=ved&password=pass",
	}
	fmt.Println(partsOfUrl.String()) // construct url from above object
}