package main

import "fmt"

func main() {

	//Predefined variable of type string 
	var fname string = "Rohan"
	//Go recognises that the variable is of type string through the value that is defined
	var lname = "Eswara"
	//In order to skip using the keyword var we can use the below shorthand
	city := "Boulder"

	fmt.Println(fname, lname, city)

	//Similarly for integer
	var age int = 26
	var birthyear = 1997
	monthnumber := 9

	fmt.Println(age, birthyear, monthnumber)

	//bits and memory
	var birthdate uint8 = 6

	//this causes an error as the range of integer values that can be assigned is -127 to 128
	var testvar uint8 = 128

	//Float variables need to have the num of btis specified
	var height float32 = 180.23

	fmt.Println(birthdate, testvar, height)
    
}
