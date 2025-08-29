package main

import (
	"fmt"
)

// slice -> dynamic array
// most used construcrt in go
// + useful methods

func main() {

	// // unintialized slice is nil
	// var nums []int
	// println(nums == nil) // true

	// fmt.Println(len(nums))

	// var nums = make([]int, 5) // len, cap
	// println(nums == nil)      // false

	// copy function

	// var nums = make([]int, 0, 5) // len, cap
	// var nums2 = make([]int, len(nums))

	// nums = append(nums, 2)

	// fmt.Println(nums, nums2)

	// Slice opearator

	// var nums = []int{1, 2, 3}

	// fmt.Println(nums[1:2])
	// fmt.Println(nums[1:])

	//slice boolean check

	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 5}

	// fmt.Println(slices.Equal(nums1, nums2)) // false)

	// 2d slice
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)

}

//run command: go run 9_Slices/slices.go
