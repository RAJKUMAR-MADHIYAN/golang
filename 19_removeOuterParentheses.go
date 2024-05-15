package main

import "fmt"

func removeOuterParentheses() {
	var S [3]string = [3]string{"()()", "(", "hari"} // S is valid parentheses string, so, s is empty "" or starts with "("
	if len(S) == 0 {
		return
	}

	var stackLength, left int
	var ret string
	for i := 0; i < len(S); i++ {
		if stackLength == 0 {
			left = i
		}
		if S[i] == "(" {
			stackLength++
			fmt.Println("hari ")
		} else {
			stackLength--
			fmt.Println("balu ")
		}
		if stackLength == 0 {
			ret += S[left+1]
			fmt.Println("raj ")

		}

	}

}
func main() {
	removeOuterParentheses()
}
