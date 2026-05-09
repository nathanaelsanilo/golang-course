package main

import (
	"fmt"
	// "strconv"
)

func greet(name string, age int) string {
	fmt.Println("hello world!")

	// answer 1
	// return "Hi, I'm " + name + " and I'm " + strconv.Itoa(age) + " years old"

	// alternative
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old", name, age)
}
