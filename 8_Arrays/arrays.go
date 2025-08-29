package main

// numbered

func main() {

	// var nums [5]int

	// nums[0] = 42
	// fmt.Println(nums[0])

	// // array length
	// fmt.Println(len(nums))

	// var vals [4]bool
	// fmt.Println(vals)

	// 2d array
	// nums := [2][2]int{{1, 2}, {3, 4}}
	// fmt.Println(nums)

}

// run with: go run 8_Arrays/arrays.go

// - fixed SIZE, THAT IS PREDICTABLE
// - MEMORY OPTIMAZIATION
// - FAST
// - COPY BY VALUE (can be changed to reference by using pointers)
// - not as flexible as slices
// - useful for MATRIX MATH, IMAGE PROCESSING, CACHING, etc.
// - can be multi-dimensional
// - cannot be resized
// - cannot use append()
// - cannot use range to add or remove elements
// - cannot use built-in functions like copy() or append()
// - can be compared using == and != operators
// - zero value is an array of zero values of the element type
// - can be initialized using a composite literal
// - can be iterated using for loop or range loop
// - can be passed to functions as arguments (by value)
// - can be returned from functions (by value)
// - can be sliced to create a slice (but the array itself is not a slice)
// - can be used as a field in a struct
// - can be used as a key in a map (if the element type is comparable)
// - can be used with the built-in functions len() and cap()
// - can be used with the built-in function copy() to copy elements from one array to another (if they have the same length and element type)
// - can be used with the built-in function append() to create a new slice from an array (but the array itself is not modified)
