package main

import "fmt"

func main() {

	//Ways to write an array
	var names = [3]string{"Rohan", "Ryan", "Ronald"}

	ages := [3]int{26, 27, 28}

	fmt.Println(names, ages)

	//Ways to write slices
	var cities = []string{"Bangalore", "Boulder", "Buenos Aires", "Bogota", "Brussells"}

	cities = append(cities, "Basel")

	subCities := cities[:3]
	subCities1 := cities[1:]

	fmt.Println(subCities, subCities1, cities)
}