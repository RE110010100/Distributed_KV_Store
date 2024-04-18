package main

import (
	"fmt"
)

func main() {

	x := 0

	// while loop equivalent
	for x < 5 {
		fmt.Println("hello world")
		x++
	}

	//for loop equivalent
	for i := 0; i <= 5; i++ {
		fmt.Println("hello world")
	}

	names := []string{"rohan", "ryan", "ronald"}
	//another way to write a for loop
	for index, value := range names {
		fmt.Println("the name in position ", index+1, " is ", value)
	}

}
