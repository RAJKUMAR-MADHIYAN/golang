package main

import (
	"fmt"
)

func mySqrt() {
	var x int

	var (
		l int
		r = x
	)

	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid == x {
			fmt.Println(mid)
		}

		if mid*mid < x {
			if (mid+1)*(mid+1) > x {
				fmt.Println(mid)
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	fmt.Println(x)
}
func main() {
	mySqrt()
}
