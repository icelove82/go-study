package main

import "fmt"

func main() {

	// Array
	var ages1 [3]int = [3]int{10, 20, 30}
	var ages2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println(ages1, len(ages1))
	fmt.Println(ages2, len(ages2))

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	fmt.Println(names, len(names))

	// Slice
	var scores = []int{100, 200, 300}
	scores[2] = 301
	scores = append(scores, 400)
	fmt.Println(scores, len(scores))

	// Slice range
	range1 := names[1:3]
	range2 := names[2:]
	range3 := names[:3]
	range4 := append(range1, "koopa")
	fmt.Println(range1, range2, range3, range4)
}
