package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	dp[0] = append(dp[0], grid[0][0])

	for i := 1; i < m; i++ {

	}

	for i := 1; i < n; i++ {

	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {

			} else {

			}
		}
	}
	fmt.Println("output")
	return n
}
func main() {
	minPathSum([][]int{})
}
