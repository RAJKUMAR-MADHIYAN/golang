package main

import "fmt"

func reverseVowels(s string) string {
	var (
		l int
		r = len(s) - 1
	)

	for r > l {

		for r >= 0 && !isVowel(s[r]) {
			r--
		}

		for l < len(s) && !isVowel(s[l]) {
			l++
		}

		if l >= r {
			break
		}

		r--
		l++
	}
	fmt.Println(s)
	return s
}
func isVowel(char byte) bool {
	vowels := [...]byte{'a', 'o', 'e', 'i', 'u', 'A', 'O', 'E', 'I', 'U'}

	for _, k := range vowels {
		if char == k {
			return true
		}
	}
	return false
}
func main() {
	reverseVowels("hello")
}
