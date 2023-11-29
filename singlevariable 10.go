package main

import (
	"fmt"
)

func hotwater(i1 int) {

	fmt.Println(i1)
}
func singlevariable() {
	fmt.Println("sigle variable")
	var i1 int = 7
	var j int
	i1 = 8
	j = i1
	fmt.Println(i1)
	hotwater(i1)
	i1 = i1 * j
	fmt.Println(i1)
	i1 = i1 % j
	fmt.Println(i1)

	for i1 <= 5 {

		fmt.Println(i1)
		hotwater(i1)
	}
}
func main() {
	singlevariable()
}
