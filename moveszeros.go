package main

import (
	"fmt"
)

func moveszeros() {
	nums := [3]int{0, 25, 35}
	var k int
	for i := range nums {

		if nums[i] != 0 {
			if k != i {
				nums[i], nums[k] = nums[k], nums[i]
				fmt.Println(nums)
			}
			k++
		}
	}

}
func main() {
	moveszeros()
}
