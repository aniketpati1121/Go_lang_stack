package main

import "fmt"

func main() {

	//const name = "Golang"
	//const age = 30

	//fmt.Println(name)
	//fmt.Println(age)

	const (
		isAdult = true
		country = "USA"
		pi      = 3.14
	)
	fmt.Println(isAdult, country, pi)

}

// const cannot be declared using := syntax
// const must be assigned a value at the time of declaration
// const cannot be reassigned a value later in the code
// const can be declared at package level or function level
// const can be of character, string, boolean, or numeric types
// const cannot be declared using the var keyword
// const can be grouped using parentheses

//go run 4_constants/con.go
