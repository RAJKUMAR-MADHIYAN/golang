package main

import "fmt"

func binarysearch() {
	var arr [4]int = [4]int{1, 2, 3, 4}
	var target int = 1
	var l int
	var r int = 4 - 1
	for l <= r {
		mid := l + (r-1)/2
		if arr[mid] == target {
			fmt.Println(mid)
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
}
func main() {
	binarysearch()
}
