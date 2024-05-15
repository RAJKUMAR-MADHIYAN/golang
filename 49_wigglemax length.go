package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	n := len(nums)

	if n == 1 || n == 0 {
		return n
	}

	up := make([]int, n)
	for i := 0; i < n; i++ {
		up[i] = 1
	}
	down := make([]int, n)
	copy(down, up)

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {

			} else if nums[i] < nums[j] {

			}
		}
	}
	fmt.Println(nums)
	return 0
}
func main() {
	wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
