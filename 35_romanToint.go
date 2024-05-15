package main

import (
	"fmt"
)

func romantoint() {
	var s string = "rajkumar"
	if 0 == len(s) {
		return
	}
	digits := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := digits[s[len(s)-1]]
	for i := len(s) - 1; i > 0; i-- {
		cur := digits[s[i]]
		pre := digits[s[i-1]]
		if cur > pre {
			fmt.Println("integer")
		} else {
			fmt.Println("nointeger")
		}
	}
	fmt.Println(sum)
}
func main() {
	romantoint()
}
