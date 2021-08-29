package main

import (
	"fmt"
	"math"
)

// Same runtime that the original solution
func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	maxProfit := maxProfit(nums)
	fmt.Println(maxProfit)
}

func maxProfit(prices []int) int {
	min, max := math.MaxInt32, 0
	for _, price := range prices {
		min = getMin(min, price)
		max = getMax(price-min, max)
	}

	return max
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
