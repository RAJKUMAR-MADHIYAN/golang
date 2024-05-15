package main

import "fmt"

func containduplicates(nums []int) bool {
	record := make(map[int]int)
	for _, num := range nums {
		if _, ok := record[num]; !ok {
			record[num] = 1
			fmt.Println("false")
		} else {
			fmt.Println("true")
		}
	}
	return false

}
func main() {
	fmt.Println(containduplicates([]int{1, 1, 2, 3, 4}))
}
