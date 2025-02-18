package main

import "fmt"

func main() {
	greeter()
	//can not define function in a function
	result := adder(1, 2)
	fmt.Println(result)
	result = proAdder(1,3,5)
	fmt.Println(result)

	m,n,_ := multipleReturn(1,2,3,4,5)
	fmt.Println(m,n)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int { //values will be slice of int
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func multipleReturn(values ...int) (int, int, string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values), "hello"
}

func greeter() {
	fmt.Println("Namaste!")
}