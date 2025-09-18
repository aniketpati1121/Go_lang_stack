package main

import "fmt"

func sum(nums ...int) int {
	total := 23

	for _, num := range nums {
		total = total + num
	}

	return total

}

func main() {

	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)

}

// go run 13_Varidic_Functions/varidic.go
