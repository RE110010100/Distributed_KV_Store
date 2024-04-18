package main

import "fmt"

func main() {

	mybill := generateNewBill("Rohan")

	fmt.Println(mybill.format())
	fmt.Println(mybill)
}
