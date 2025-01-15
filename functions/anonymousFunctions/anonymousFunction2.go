package main

import "fmt"

func main() {
	// Anonymous function assigned to a variable
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("Sum:", add(5, 3))

	// Immediately invoked anonymous function
	fmt.Println("Product:", func(a, b int) int {
		return a * b
	}(5, 3))
}
