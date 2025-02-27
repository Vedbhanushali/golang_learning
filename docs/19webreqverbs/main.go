package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformGetRequest()
	PerformPostJSONRequest()
	PerformPostFormRequest()
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

func PerformPostJSONRequest() {
	const myurl = "https://localhost:8080/post"

	// fake json payload
	requestBody := strings.NewReader(`{
		"key1": "value1",
		"key2": "value2"
	}`)
	response, error := http.Post(myurl, "application/json", requestBody)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "https://localhost:8080/postjson"

	// fake form data
	data := url.Values{}
	data.Add("firstname", "John")
	data.Add("lastname", "Doe")
	data.Add("email", "test@go.dev")

	respose , err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer respose.Body.Close()
	content, _ := io.ReadAll(respose.Body)
	fmt.Println(string(content))
}