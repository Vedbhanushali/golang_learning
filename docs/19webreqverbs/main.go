package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://localhost:8080/get"

	response, error := http.Get(myurl)

	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	fmt.Println("status ",response.Status)
	fmt.Println("content length",response.ContentLength)

	//another way to create string from response body
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("byte count",byteCount)
	fmt.Println("response body",responseString.String())
}