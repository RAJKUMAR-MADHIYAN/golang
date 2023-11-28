package main

import (
	"fmt"
)

func singlenumber() {
	nums := []int{1, 2, 3}

	res := 3
	for _, num := range nums {
		res ^= num
	}
	fmt.Println(res)

}
func main() {
	singlenumber()
}
