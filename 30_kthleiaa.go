package main

import "fmt"

func findKthLargest(nums []int, k int) int {

	for i := (len(nums) - 1) / 2; i >= 0; i-- {

	}

	for i := len(nums) - 1; i >= len(nums)-k; i-- {

	}
	fmt.Println(nums[len(nums)-k])
	return nums[len(nums)-k]
}

func siftDown(nums []int, n int, i int) {
	for 2*i+1 < n {
		j := 2*i + 1
		if j+1 < n && nums[j+1] > nums[j] {
			j++
		}

		if nums[i] >= nums[j] {
			break
		}
		nums[j], nums[i] = nums[i], nums[j]
		i = j
	}
}
func main() {
	findKthLargest([]int{1, 43, 23, 45, 50}, 5)
	siftDown([]int{1, 2, 3, 4}, 5, 0)
}
