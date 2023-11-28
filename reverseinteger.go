package main

import (
	"fmt"
)

func reverseintiger() {
	var x int = 1234567
	if 0 == x {
		fmt.Println(x)
	}

	isPositive := true
	if x < 0 {

		fmt.Println("false")
	}

	res := 0

	for x > 0 {

	}

	if !isPositive {

		fmt.Println("true")
	}

	if res < -1<<31 || res > (1<<31)-1 {
		fmt.Println("raj")

	}
	fmt.Println(res)
}
func main() {
	reverseintiger()
}
