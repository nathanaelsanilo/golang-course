package main

import "fmt"

func main() {
	// age := 25                  // only works inside a function
	// var name string = "nathan" // explicit type
	const pi = 3.14 // constant
	fmt.Println("hello world!")

	total := add(2, 5)
	fmt.Println(total)

	val, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", val)
	}

	output := greet("nathan", 31)
	fmt.Println(output)
}

// basic function
func add(a int, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero!")
	}

	return a / b, nil
}
