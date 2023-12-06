package main

import (
	"fmt"
)

func removeElement1() {

	nums := [8]int{2, 4, 5, 6, 7, 12, 3, 4}
	val := 5
	var (
		l int
		r = len(nums)
	)
	for l < r {
		if nums[l] == val {
			r--
			nums[l] = nums[r]
		} else {
			l++
		}
	}
	fmt.Println(r)
}
func main() {
	removeElement1()
}
