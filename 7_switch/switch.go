package main

func main() {

	// simple switch statement

	// i := 2

	// switch i {
	// case 1:
	// 	println("one")
	// case 2:
	// 	println("two")
	// case 3:
	// 	println("three")
	// case 4:
	// 	println("four")
	// case 5:
	// 	println("five")
	// default:
	// 	println("unknown number")
	// }

	// Muliple condition switch statement

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("It's the weekend!")
	// default:
	// 	println("It's a weekday.")
	// }

	// type switch statement

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			println("I'm an int")
		case string:
			println("I'm a string")
		case bool:
			println("I'm a bool")
		default:
			println("Unknown type", v)
		}

	}
	whoAmI(42)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(3.14)

}

// run command : go run 7_switch/switch.go
