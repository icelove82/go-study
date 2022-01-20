package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is: ", i)
	}

	names := []string{"ymx", "kim", "yang"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, v := range names {
		fmt.Printf("the value in index %v is %v. \n", i, v)
	}
}
