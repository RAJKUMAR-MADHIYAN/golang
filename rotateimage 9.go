package main

import (
	"fmt"
)

func rotateimage() {
	matrix := [][]int{{1, 2},
		{2, 3},
	}
	n := len(matrix)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			fmt.Println("rotate image")

		}
	}

}
func main() {
	rotateimage()
}
