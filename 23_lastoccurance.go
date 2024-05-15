package main

import "fmt"

func lastOccurance(nums []int, target int) int {
	var (
		l int
		r = len(nums)
	)
	for l != r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println("last potion of sorted array", l)
	return l
}
func main() {
	lastOccurance([]int{1, 2, 3, 4, 5}, 8)
}
