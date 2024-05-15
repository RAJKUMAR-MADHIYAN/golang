package main

import "fmt"

func climbStairs(n int) (int, []int) {
	var memo = []int{1, 1}
	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}

	return memo[n], memo
}
func main() {
	fmt.Println(climbStairs(5))
}
