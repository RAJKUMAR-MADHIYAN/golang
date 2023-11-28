package main

import (
	"fmt"
)

func maxprofit() {
	var prices [3]int = [3]int{2, 3, 6}
	n := len(prices)
	if 0 == n || 1 == n {

	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > res {
				res = prices[i] - prices[j]
			}
		}
	}
	fmt.Println(res)

}
func main() {
	maxprofit()
}
