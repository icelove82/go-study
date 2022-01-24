package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {

	firstName, lastName := getInitials("tifa lockhart")
	fmt.Println(firstName, lastName)

	f2, l2 := getInitials("Alnord YUN")
	fmt.Println(f2, l2)

	f3, l3 := getInitials("Tom")
	fmt.Println(f3, l3)

}
