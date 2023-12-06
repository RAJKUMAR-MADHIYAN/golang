package main

import "fmt"

func intersection() {

	nums1 := [3]int{1, 2, 3}
	nums2 := [3]int{4, 5, 6}
	set1 := nums1.NewSet()
	for _, num := range nums1 {
		set1.Add(num)
	}
	set2 := nums2.NewSet()
	for _, num := range nums2 {
		set2.Add(num)
	}

	var res []int
	for item := range set1 {
		if set2.Contains(item) {
			if value, ok := item.(int); ok {
				res = append(res, value)
			}
		}
	}
	fmt.Println(res)

}
func main() {
	intersection()
}
