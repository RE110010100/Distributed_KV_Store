package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	names := strings.Split(n, " ")
	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) >= 2 {
		return initials[0], initials[1]
	}

	return initials[0], "-"

}

func main() {

	fn, sn := getInitials("Rohan Eswara")
	fn1, sn1 := getInitials("Jackie")

	fmt.Println(fn, sn)
	fmt.Println(fn1, sn1)
}
