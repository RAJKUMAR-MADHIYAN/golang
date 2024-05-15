package main

import "fmt"

func plusOne(digits []int) []int {
	var (
		carry = 1
		index = len(digits) - 1
	)
	for index >= 0 {
		sum := carry + digits[index]
		digits[index] = sum % 10
		carry = sum / 10
		index--
	}

	if carry > 0 {
		digits = append(digits, 0)
		copy(digits[1:], digits[0:])
		digits[0] = carry
	}
	fmt.Println(digits)
	return digits
}
func main() {
	plusOne([]int{1, 2, 3, 4, 5})
}
