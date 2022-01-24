package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// loop in map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		9989: "mario",
		5918: "luigi",
		0432: "peach",
	}
	phonebook[2116] = "bowser"
	phonebook[9989] = "mario-2"

	fmt.Println(phonebook)
	fmt.Println(phonebook[5918])
}
