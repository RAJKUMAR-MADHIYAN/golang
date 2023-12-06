package main

import "fmt"

func reversestring() {
	var s string = "rajkumar"
	sLen := len(s)
	bytes := []byte(s)
	for i := 0; i < sLen/2; i++ {
		bytes[i], bytes[sLen-i-1] = bytes[sLen-i-1], bytes[i]
	}
	fmt.Println(bytes)

}
func main() {
	reversestring()
}
