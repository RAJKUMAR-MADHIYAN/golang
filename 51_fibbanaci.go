package main

import "fmt"

func fibbni() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}

}
func main() {
	n := fibbni()
	for i := 0; i < 10; i++ {
		fmt.Println(n())
	}
}
