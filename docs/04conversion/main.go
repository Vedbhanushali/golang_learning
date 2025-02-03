package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	num,error := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("+2 is",  num + 2)
		fmt.Printf("Type of input is %T %T", num, input)
	}
}