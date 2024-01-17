package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5) // remove {}

	userNames[0] = "Julie"

	userNames = append(userNames, "Loai")
	userNames = append(userNames, "Al-Sharee")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3) // remove {}

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9

	fmt.Println(courseRatings)
}
