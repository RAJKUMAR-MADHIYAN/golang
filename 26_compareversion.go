package main

import (
	"fmt"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	var (
		v1 = strings.Split(version1, ".")
		v2 = strings.Split(version2, ".")
	)
	for i := 0; i < len(v1) || i < len(v2); i++ {
		var v1N, v2N int
		if i < len(v1) {
			fmt.Println("v21")
		}
		if i < len(v2) {
			fmt.Println("v12")
		}

		if v1N > v2N {
			return 1
		} else if v1N < v2N {
			return -1
		}
	}
	return 0
}
func main() {
	compareVersion("rajkumar", "balamurugan")

}
