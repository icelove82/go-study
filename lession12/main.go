package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func updateMenu(x map[string]float64) {
	x["coffee"] = 5.55
}

func main() {

	// Type group A -> strings, ints, bools, floats, arrays, structs
	// Value pass
	name := "tifa"

	updateName(name)

	fmt.Println(name)

	// Pointer
	fmt.Println("Memory address : ", &name)
	m := &name
	fmt.Println("Value :", *m)

	// Type group B -> slices, maps, functions
	// Reference pass
	menu := map[string]float64{
		"coffee": 4.99,
		"pie":    7.99,
	}

	updateMenu(menu)

	fmt.Println(menu)
	fmt.Println(menu["coffee"])
}
