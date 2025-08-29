package main

import "fmt"

// maps -> hash,object,dict

func main() {
	// //creating map
	// m := make(map[string]string)

	// //setting an element

	// m["name"] = "John"
	// m["age"] = "30"
	// m["country"] = "USA"

	// //getting an element
	// fmt.Println(m["name"], m["age"], m["country"])

	m := make(map[string]int)
	m["a"] = 1

	fmt.Println(m["a"])
}
