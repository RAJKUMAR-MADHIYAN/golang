package main

import (
	"fmt"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)

	for num := left; num <= right; num++ {
		flag := true
		strNum := strconv.Itoa(num)

		for j := 0; j < len(strNum); j++ {
			if divisor := int(strNum[j] - '0'); divisor == 0 || num%divisor != 0 {
				flag = false
			}
		}
		if flag {
			res = append(res, num)
		}
	}
	fmt.Println(res)
	return res
}
func main() {
	selfDividingNumbers(4, 5)
}
