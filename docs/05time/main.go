package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time line")
	presentTime := time.Now()
	fmt.Println("current time is", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Mon Monday")) // this is standard syntax for formating

	createdDate := time.Date(2025, time.April, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("created date is", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Mon Monday"))
}