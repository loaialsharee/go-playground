package main

import "fmt"

func main() {
	prices := [4]float64{10, 11, 12, 13.2}
	fmt.Println(prices)

	// managing slices
	priceList := []float64{10, 8}
	fmt.Println(prices[0:1])
	priceList[1] = 9.99

	// append
	priceList = append(priceList, 5.99)
	priceList = priceList[1:] // to remove the last element of the slice
	fmt.Println(priceList)

	// appending slice to another slice
	nums := []float64{5, 12, 16, 20}
	nums = nums[1:]
	fmt.Println(nums)

	discounts := []float64{1, 2, 3, 4}
	nums = append(nums, discounts...) // three dots to decomposte them and append per item
}
