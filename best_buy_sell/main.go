package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	maxProfit := maxProfit(nums)
	fmt.Println(maxProfit)
}

func maxProfit(prices []int) int {
	minSeen := prices[0]
	maxSeen := 0

	if len(prices) == 0 {
		return 0
	}

	// On golang, when using a for range the first is the index, and the second is a copy of the element of that index. https://tour.golang.org/moretypes/16
	for _, currentPrice := range prices {
		if currentPrice < minSeen {
			minSeen = currentPrice
		} else if (currentPrice - minSeen) > maxSeen {
			maxSeen = currentPrice - minSeen
		}
	}

	return maxSeen
}
