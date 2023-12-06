package main

import (
	"fmt"
	"math"
)

func maxsubarray() {
	var nums [4]int = [4]int{3, 2, -5, 4}
	var (
		max           = math.MinInt32
		sum           int
		start, end, f int
	)
	for i, j := range nums {
		sum += j
		if sum > max {
			if f != 0 {
				start = f + 1
				f = 0
			}
			max = sum
			end = i
		}
		if sum < 0 {
			
		}
	}
	fmt.Println("start index is %d, end index is %d\n", start, end)
	fmt.Println(max)

}
func main() {
	maxsubarray()
}
