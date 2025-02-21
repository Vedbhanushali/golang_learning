package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	

	content := "Hello"

	file, error := os.Create("./file.txt")
	checkNilErr(error)

	length, err := io.WriteString(file, content)
	// if error != nil {
	// 	panic(error) //will shutdown the program and print the error message
	// 	// just like throw
	// } //moved this code to below func
	checkNilErr(err)
	fmt.Println("Length: ", length)
	defer file.Close() //mean at end of function this will be executed in case need to perform some code after file closing
	readFile("./file.txt")
}

func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Data: ", databytes,string(databytes))
}

func checkNilErr(err error) {
	if error != nil {
		panic(error) //will shutdown the program and print the error message
		// just like throw
	}
}