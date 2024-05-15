package main

import "fmt"

func addDigits(num int) int {

	for num > 9 {
		num = performAdd(num)
	}
	fmt.Println(num,)
	return num
}

func performAdd(num int) (res int) {
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return
}
func main() {
	addDigits(5)
}
