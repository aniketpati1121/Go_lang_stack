package main

import "fmt"

func counter() func() int {

	var count int = 9

	return func() int {
		count += 5
		return count
	}

}

func main() {

	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

// go run 14_Closures/closures.go
