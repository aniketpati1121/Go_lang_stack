package main

import "fmt"

// iterating over a data structure is called range
func main() {

	nums := []int{1, 2, 3, 4, 5}

	sum := 0

	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println((sum))

	// for i := 0; i < len(nums); i++ {
	// 	println(nums[i])
	// }

}

// go run 11_range/range.go
