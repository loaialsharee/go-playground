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
		fmt.Println("Fan In process")
	default:
		fmt.Println("Invalid process")
	}
}
