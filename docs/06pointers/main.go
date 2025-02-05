package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")
	var ptr *int //* for pointer declaration
	fmt.Println("value of ptr is",ptr)

	myNum := 10
	var numPtr *int = &myNum // & for memory address of variable is pointing to by reference
	fmt.Println("value of numPtr is",numPtr,*numPtr)
	fmt.Println("value of myNum is",myNum)
	*numPtr = *numPtr * 2 //memory address of myNum is modified
	fmt.Println("value of myNum is",myNum)

}