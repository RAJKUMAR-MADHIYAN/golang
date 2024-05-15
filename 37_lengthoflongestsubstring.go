package main

import "fmt"

func lengthoflongestsubstring(s string) int {
	var (
		start      = 0
		res        = 0
		lastrecord = make(map[rune]int)
	)
	for i, ch := range []rune(s) {
		if lastindex, ok := lastrecord[ch]; ok && lastindex >= start {
			start = lastindex + 1
		}
		if i-start+1 > res {
			res = i - start + 1
		}
		lastrecord[ch] = i

	}
	fmt.Println(res)
	return res
}
func main() {
	lengthoflongestsubstring("hello world")
}
