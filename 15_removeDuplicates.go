package main

import "fmt"

func nextDifferentCharacterIndex(nums [5]int, p int) int {

	//var nums [3]int = [3]int{1, 2, 1}
	for ; p < len(nums); p++ {
		if nums[p] != nums[p-1] {
			break
		}
	}

	return p
}

func removeDuplicates() {
	var nums [5]int = [5]int{1, 0, 5, 1, 5}
	n := len(nums)

	if 0 == n {
		fmt.Println(n)
	}

	var (
		res   = 1
		i     = 1
		index = nextDifferentCharacterIndex(nums, 1)
	)

	for index < n {
		res++
		nums[i] = nums[index]
		i++
		index = nextDifferentCharacterIndex(nums, index+1)
	}
	fmt.Println(res)
}
func main() {
	removeDuplicates()
}
