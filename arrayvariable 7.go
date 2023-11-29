package main

import (
	"fmt"
)

func arraycopy(arr [3]int) {
	for i1 := 0; i1 < 3; i1++ {
		for j := 0; j < i1; j++ {
			fmt.Println("inner loop")
		}
	}

}
func arrayvariable() {
	fmt.Println("array variable")
	var arr [3]int = [3]int{1, 2, 3}
	arr[0] = 5
	arr[1] = 10
	arr[2] = 20
	arr[0] = arr[1]
	fmt.Println(arr)
	arraycopy(arr)
	arr[1] = arr[0] + arr[1]
	arr[0] = arr[1] + arr[2]
	for i1 := 0; i1 < 3; i1++ {
		j := i1
		if j == i1 {

		}
		if arr[i1] > arr[i1]+1 {

		}

	}

}
func main() {
	arrayvariable()
}
