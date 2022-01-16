package main

import "fmt"

func main() {

	// string
	var name1 string = "mario"
	var name2 = "luigi"
	var name3 string

	fmt.Println(name1, name2, name3)

	name1 = "peach"
	name3 = "bowser"

	fmt.Println(name1, name2, name3)

	name4 := "yoshi"
	fmt.Println(name4)

	// int
	var age1 int = 20
	var age2 = 30
	age3 := 40

	fmt.Println(age1, age2, age3)

	// bit & memory
	var num1 int8 = 25
	var num2 int8 = -125
	var num3 uint16 = 256

	fmt.Println(num1, num2, num3)

	// float
	var score1 float32 = 25.98
	var score2 float64 = 83928439084923849.9
	score3 := 1.5

	fmt.Println(score1, score2, score3)
}
