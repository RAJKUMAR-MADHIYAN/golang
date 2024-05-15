package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if 1 == n {
		return true
	}

	cur := 0
	flag := false

	for cur < n {
		if 0 == bits[cur] {
			cur++
		} else {
			cur += 2
		}
		if cur == n-1 {
			flag = true
		}
	}
	fmt.Println(flag)
	return flag
}
func main() {
	isOneBitCharacter([]int{1, 2, 3, 4})
}
