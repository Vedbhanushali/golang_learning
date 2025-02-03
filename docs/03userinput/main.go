package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://pkg.go.dev/ - will import buffer packages to read IO operations
func main() {
	var welcome string = "Welcome to user input program."
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name : ")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n') // \n is delimiter stops reading the 
	// input, _ := reader.ReadString('\n') // _ is used to ignore the error variable if not going to use it
	// _, error //to ignore input and use only error
	fmt.Println("Welcome ",input)
	fmt.Printf("Type of input is %T", input)

	i1, error := reader.ReadString('\n')
	if error != nil {
		fmt.Println("An error occured while reading input. Please try again.")
	} else {
		fmt.Println("Welcome ",i1)
		fmt.Printf("Type of input is %T", i1)
	}
}