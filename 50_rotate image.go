package main

import "fmt"

func rotateimage() {
	matrix := [][]int{{1, 2},
		{2, 3},
	}
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			fmt.Println("rotate image", matrix)

		}
	}

}
func main() {
	rotateimage()
}
