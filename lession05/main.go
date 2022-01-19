package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// strings package
	greeting := "Hello there friends!"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// sort packages
	ages := []int{45, 2, 25, 54, 42, 6, 76, 34, 15, 63}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 1)
	fmt.Println(index)

	names := []string{"ymx", "kim", "yang"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "kim"))
}
