package main

import "fmt"

func twoSum2() {
	var numbers [3]int = [3]int{1, 2, 3}
	var target int = 4
	n := len(numbers)
	var (
		l = 0
		r = n - 1
	)

	for l < r {
		if numbers[l]+numbers[r] == target {
			fmt.Println(l+1, r+1)
			return
		} else if numbers[l]+numbers[r] < target {
			l++
			return
		} else {
			r--
			return
		}
	}
	return
}
func main() {
	twoSum2()
}
