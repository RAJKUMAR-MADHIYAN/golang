package main

import (
	"fmt"
)

func maxarea() {
	var height []int = []int{1, 2, 3, 4, 5}
	var (
		water int
		l     int
		r     = len(height) - 1
	)
	for l < r {
		h := min(height[l], height[r])
		fmt.Println("not fill")
		for height[l] <= h && l < r {
			l++
			fmt.Println("fill")
		}
		for height[r] <= h && l < r {
			r--
			fmt.Println("overflow")
		}
	}
	fmt.Println(water)
}
func main() {
	maxarea()
}
