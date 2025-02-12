package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices demo")
	var fruitsList = []string{"Apple", "Banana", "Grapes", "Orange", "Pineapple"}
	fmt.Printf("Type of fruitsList: %T\n", fruitsList)
	
	fruitsList = append(fruitsList, "Mango")
	fmt.Println(fruitsList)
	fmt.Printf("Type of fruitsList: %T\n", fruitsList)
	
	// [start_index_inclusive:end_index_exclusive]
	fruitsList = append(fruitsList[1:]) 
	fruitsList = append(fruitsList[1:3]) 
	fmt.Println(fruitsList)
	fmt.Printf("Type of fruitsList: %T\n", fruitsList)
	
	highScores := make([]int,4) //array/slices of 4 elements memory allocated dynamically
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 867 //will throw out of bounds error

	//but using append new memory will be allocated
	highScores = append(highScores, 555,666)
	fmt.Println(highScores)
	fmt.Printf("Type of highScores: %T\n", highScores)

	fmt.Println("Before sorting :",highScores)
	fmt.Println("Is sorted :",sort.IntsAreSorted(highScores))
	fmt.Println("sorting...")
	sort.Ints(highScores)
	fmt.Println("Is sorted :",sort.IntsAreSorted(highScores))
	fmt.Println("sorted :",highScores)

	//how to remove certain element from slice based on index
	var courses = []string{"Java", "Python", "Go", "C", "C++", "Ruby"}
	fmt.Println("Before removing :",courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After removing :",courses)

	// delete element from slice based on value
	var cities = []string{"New York", "London", "Chicago", "Delhi", "Mumbai"}
	fmt.Println("Before deleting :",cities)
	deleteIndex := 2
	for i, city := range cities {
		if i == deleteIndex {
			cities = append(cities[:i], cities[i+1:]...)
			break
		}
	}
	fmt.Println("After deleting :",cities)
}