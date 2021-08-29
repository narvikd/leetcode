package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Println(result)
}

// twoSums has a time and space complexity: O(n), n is length of the array.
func twoSum(nums []int, target int) []int {
	// 1.Create a hashmap
	seen := make(map[int]int)

	// 2. Iterate the nums.
	//	  Remember: On golang, when using a for range the first is the index, and the second is a copy of the element of that index. https://tour.golang.org/moretypes/16
	//	  In other words, nums[i] = x
	for i, x := range nums {
		// 3. Check if the number is in the map. https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
		//	  and https://golang.org/doc/effective_go#maps -> ctrl + f "To test for presence in the map without worrying about the actual value"
		_, found := seen[x]
		if found {
			return []int{seen[x], i}
		}

		// 3. Add the complementary number from the target to the hashmap, in order to access it later, that way it only has to iterate once.
		seen[target-x] = i
	}
	// Example:
	// If [2, 7, 11, 15] & target = 9
	// At the first iteration x(nums[i]) will be 2 at index(i) 0.
	// What will be stored in the hashmap will be 7.
	// target - x = 7.		9 - 2 = 7
	// seen[7] = 0
	//
	// On the next iteration (x=7): The runtime will ask, have we seen a 7?, then it returns true and executes the if-return statement.
	return nil
}
