package main

import "fmt"

//struct are kind of classes in go
func main() {
	fmt.Println("Structs in Go")
	//no inheritance in golang, no super or parent

	ved := User{"Ved", "ved@go.dev", true, 21}
	fmt.Println(ved, ved.Name)
	fmt.Printf("User Ved: %+v\n", ved) // %+v will print the field names as well

	fmt.Println("Is Ved Active? ", ved.GetStatus())
	ved.NewMail() //original email does not change pass by value not reference
	fmt.Println("Email of this user is: ", ved.Email) 
	ved.NewMail2()
	fmt.Println("Email of this user is: ", ved.Email) 
}

type User struct { 
	Name string
	Email string
	Status bool
	Age int 
}

func (u User) GetStatus() bool {
	return u.Status
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

func (u *User) NewMail2() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

