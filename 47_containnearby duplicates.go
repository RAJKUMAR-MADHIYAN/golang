package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := record[v]; ok {
			if i-j <= k {
				fmt.Println("true")
				return true
			}
		}
		record[v] = i
	}
	fmt.Println(false)
	return false
}
func main() {
	containsNearbyDuplicate([]int{1, 2, 3, 4, 5}, 5)
}
