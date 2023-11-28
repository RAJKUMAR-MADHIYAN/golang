package main

import (
	"fmt"
)

func twoSum1() {
	nums := [5]int{1, 2, 3, 4, 5}
	target := 7
	for i, j := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k]+j == target {
				fmt.Println(nums[i], nums[k])
			}
		}
	}
	fmt.Println(nums)
}
func main() {
	twoSum1()
}
