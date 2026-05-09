package main

import "fmt"

func main() {
	var process string
	fmt.Print("Enter the process key: ")
	fmt.Scan(&process)

	switch process {
	case "fanOut":
		processFanOut()
	case "fanIn":
		processFanIn()
	default:
		fmt.Println("Invalid process")
	}
}
