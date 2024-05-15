package main

import "fmt"

func searchinsert(nums [5]int, target int) int {
	var (
		l int
		r = len(nums) - 1
	)
	for l <= r {
		mid := l + (r-1)/2
		if target == nums[mid] {
			fmt.Println("if1", len(nums), l)
			return mid
		}
		if target < nums[mid] {

			r = mid - 1
			fmt.Println("if2", l)
		} else {

			l = mid + 1
			fmt.Println("if3", l)
		}
	}
	fmt.Println(l)
	return l
}
func main() {
	searchinsert([5]int{1, 2, 3, 4, 5}, 4)
}
