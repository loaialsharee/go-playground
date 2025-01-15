package main

import "fmt"

func main() {
	// Closure capturing the variable `count`
	counter := func() func() int {
		count := 0 // Enclosed variable
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("Counter:", counter()) // Output: Counter: 1
	fmt.Println("Counter:", counter()) // Output: Counter: 2
	fmt.Println("Counter:", counter()) // Output: Counter: 3
}
