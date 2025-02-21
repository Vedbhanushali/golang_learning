package main

import "fmt"

func main() {
	// firstExample()
	// it will print three two one where poped from the stack
	// it follows the LIFO order STACK so stack is left is bottom and right is top | one | two | three |
	defer fmt.Println("one")  
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello")
	defer fmt.Println("four")
	myDefer()
}

func myDefer() {
	// will push in stack and popped make it reverse
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
func firstExample() {
	defer fmt.Println("world") // it is going to move to the end of the function 
	fmt.Println("hello")
}