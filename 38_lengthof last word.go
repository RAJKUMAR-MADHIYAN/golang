package main

import "fmt"

func lengthoflastword(s string) int {
	var (
		n   = len(s)
		cur = n - 1
	)
	for cur >= 0 {
		if n-1 == cur && s[cur] == 32 {
			cur--
			n--
			fmt.Println("raj")
			continue
		}
		if s[cur] != 32 {
			cur--

		} else {
			break
		}
	}
	fmt.Println((n - cur - 1))
	return n - cur - 1
}
func main() {
	lengthoflastword("rajkumar")
}
