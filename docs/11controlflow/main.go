package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If else")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)


	if 5%2 == 0 {
		fmt.Println("5 is even")
	} else {
		fmt.Println("5 is odd")
	}

	if num := 9; num < 0 { //initialization and checking in same line this variable is in local scope of if only 
		// can be used in web request response if want to use in local scope
		fmt.Println("Number is negative")
	}

	//switch case
	fmt.Println("Switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1 and you can open")
	case 2:
		fmt.Println("Dice number is 2 and you can move 2 steps")
	case 3:
		fmt.Println("Dice number is 3 and you can move 3 steps")
		//if any case hit it will break
	case 4:
		fmt.Println("Dice number is 4 and you can move 4 steps")
		fallthrough //fallthrough will execute the next case as well in case don't want to break
	case 5:
		fmt.Println("Dice number is 5 and you can move 5 steps")
	case 6:
		fmt.Println("Dice number is 6 and you can move 6 steps and roll the dice again")
	default:
		fmt.Println("Something went wrong!")
	}
}