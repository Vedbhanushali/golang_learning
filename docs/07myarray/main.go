// Most cases we use slices instead of arrays in Go. Slices are more flexible, powerful and convenient than arrays.

package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")
	var fruitsList [4]string //here size of array should be defined at the time of declaration
	fruitsList[0] = "Apple"
	fruitsList[1] = "Banana"
	fruitsList[3] = "Orange"

	fmt.Println("Fruits List is", fruitsList, len(fruitsList))

	var vegList = [3]string{"Potato", "Beans", "Brinjal"} //array initialization
	fmt.Println("Veg List is", vegList, len(vegList))
}