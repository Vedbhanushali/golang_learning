package main

import "fmt"

const BaseURL string = "https://api.com" //constant variable
//here intial Letter Capital means it is exported as public and can be used outside the package

func main() {
	var username string = "ved"
	fmt.Println("Username: ", username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("Is Logged In: ", isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallVal uint = 255 // mulitliple types of uint like uint8, uint16, uint32, uint64
	//uint storage is [0,...,255] 
	fmt.Println("Small Value: ", smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.455
	fmt.Println("Small Float Value: ", smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	//default values and some aliases
	var anotherValue int
	fmt.Println(anotherValue) //default value is 0 no garbage value
	fmt.Printf("variable is of type : %T \n", anotherValue)

	var emptyString string
	fmt.Println(emptyString) //default value is empty string
	fmt.Printf("variable is of type : %T \n", emptyString)

	// implicit type
	var website = "vedbhanushali.site" //automatically detects the type
	fmt.Println(website)

	// no var style variable declaration
	numberOfUsers := 3000 // this method can only be used inside a function
	fmt.Println(numberOfUsers)

}