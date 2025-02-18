package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	for d:=0; d<len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days { // i is index
		fmt.Println(days[i])
	}

	// for each loop
	for index, day := range days { //index value can use as _ if not required for index
		fmt.Println(index, day)
	}

	// while loop
	start := 0
	for start < len(days) {
		if(start == 2) {
			fmt.Println("Tuesday interesting")
			start++
			continue // can break also
		}
		fmt.Println(days[start])
		start++;
	}
}