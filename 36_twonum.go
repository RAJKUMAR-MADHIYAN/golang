package main

import "fmt"

func twonum(num []int, target int) []int {
	var record map[int]int = make(map[int]int)

	for i, j := range num {
		total := target - j
		if res, ok := record[total]; ok {
			fmt.Println(res, i)
			return []int{res, i}

		}
		record[j] = i
		fmt.Println(num[i], num[j], record, total)
	}
	return []int{}
}

func main() {
	twonum([]int{1, 2, 3, 4}, 6)
}
