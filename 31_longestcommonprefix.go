package main

import (
	"fmt"
)

func prefix() {
	var strs [3]string = [3]string{"flower", "flow", "float"}
	if 0 == len(strs) {
		return
	}
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				fmt.Println(res)

			}
		}
		res += string(c)
	}

	fmt.Println(prefix)
}
func main() {
	prefix()
}
