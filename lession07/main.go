package main

import "fmt"

func main() {
	age := 25

	fmt.Sprintln(age <= 50)
	fmt.Sprintln(age != 40)

	if age < 30 {
		fmt.Sprintln("age is less than 30")
	} else if age < 40 {
		fmt.Sprintln("age is less than 40")
	} else {
		fmt.Sprintln("age is not less than 45")
	}

	names := []string{"ymx", "kim", "yang", "guang", "gang"}

	for i, v := range names {
		if i == 1 {
			fmt.Printf("continue at pos %v \n", i)
			continue
		}

		if i > 2 {
			fmt.Printf("break at pos %v \n", i)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", i, v)
	}

}
