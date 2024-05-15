package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(g)))

	var (
		si, gi, res int
	)

	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			gi++
			si++
		} else {
			gi++
		}
	}
	fmt.Println(res)
	return res
}
func main() {
	findContentChildren([]int{1, 2, 3}, []int{1, 2, 3})
}
