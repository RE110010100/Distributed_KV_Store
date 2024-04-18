package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Create a new Bill, please mention the person's name : ")
	name, _ := reader.ReadString('\n')

	fmt.Println("please mention the person's age : ")
	age, _ := reader.ReadString('\n')

	ageFloat, _ := strconv.ParseFloat(age, 64)
	fmt.Print("name : ",name " age : ",ageFloat)
}
