package main

import "fmt"

func isOneBitCharacter() bool {
	var bits []int = []int{1, 2, 3, 4, 5}
	n := len(bits)
	if 1 == n {
		fmt.Println(true)

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
	isOneBitCharacter()
}
