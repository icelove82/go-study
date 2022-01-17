package main

import "fmt"

func main() {

	age := 35
	name := "shau"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("Hello YUN!")
	fmt.Println("You will success!")
	fmt.Println("Name is", name, ", Age is", age)

	// Printf (formatted strings)
	fmt.Printf("Name is %v , Age is %v \n", name, age)
	fmt.Printf("Age is of type of %T \n", age)
	fmt.Printf("The price is %0.2f \n", 15.999)

	// Save printf
	var str = fmt.Sprintf("Name is %v , Age is %v \n", name, age)
	fmt.Println("The saved string is -> ", str)

}
