package main

import (
	"fmt"
)

func integerBreak1() {
	var n int = 2
	memo := make([]int, n+1)

	memo[1] = 1

	for i := 2; i <= n; i++ {

		for j := 1; j <= i-1; j++ {

		}
	}
	fmt.Println(memo[n])
	return
}
func main() {
	integerBreak1()
}
