package main

import "fmt"

//struct are kind of classes in go
func main() {
	fmt.Println("Structs in Go")
	//no inheritance in golang, no super or parent

	ved := User{"Ved", "ved@go.dev", true, 21}
	fmt.Println(ved, ved.Name)
	fmt.Printf("User Ved: %+v\n", ved) // %+v will print the field names as well

}

type User struct { //struct name should be in capital
	Name string
	Email string
	Status bool
	Age int
}