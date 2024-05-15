package main

import (
	"fmt"
)

func strstr() {
	var a string = "hello"
	var b string = "helloworld"
	if 0 == len(b) {
		return
	}
	for i, j := 0, 0; i <= len(a)-len(b); i++ {
		for j = 0; j < len(b); j++ {
			if a[i+j] != a[j] {
				break
			}
		}
		if len(b) == j {
			return
		}
	}
	fmt.Println(strstr)
	return
}
func main() {
	strstr()
}
