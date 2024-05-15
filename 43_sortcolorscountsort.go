package main

import "fmt"

func sortColorsCountSort(nums []int) []int {
	var count = [3]int{}

	for _, num := range nums {
		count[num]++
	}

	index := 0

	for i, j := range count {
		for k := 0; k < j; k++ {
			nums[index] = i
			index++
			fmt.Println("sort count")
		}
	}
	return nums
}
func main() {
	sortColorsCountSort([]int{1, 2, 3})
}
