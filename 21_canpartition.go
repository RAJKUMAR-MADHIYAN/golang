package main

import "fmt"

func canPartition(nums []int) bool {
	var sum int = 2
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	c := sum / 2
	n := len(nums)
	memo := make([]bool, c+1)

	for i := 0; i <= c; i++ {
		memo[i] = nums[0] == i
	}

	for i := 0; i < n; i++ {
		for j := c; j >= nums[i]; j-- {
			memo[j] = memo[j] || memo[j-nums[i]]
		}
	}
	fmt.Println(memo[c])
	return memo[c]
}
func main() {
	canPartition([]int{1, 2, 3, 4, 5})
}
