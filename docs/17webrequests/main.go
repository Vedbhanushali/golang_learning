package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev"

func main() {
	response,err := http.Get(url) // get request to the url
	
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //very important to close the response body

	fmt.Printf("type of response %T\n", response)
	fmt.Printf("response status code %v\n", response.StatusCode)

	databytes , err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response body %v\n", string(databytes))
}