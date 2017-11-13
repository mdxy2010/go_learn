package main

import "fmt"

//func add() (v1 int, v2 in) {
	//return v1 + v2
//}

func main() {
	v1 := 2
	v2 := 13
	var sum int
	var dec int
	var xxx int
	var div int
	var __div int

	sum = v1 + v2
	dec = v1 - v2
	xxx = v1 * v2
	div = v2 / v1
	__div = v2 % v1

	fmt.Println("sum = ", sum)
	fmt.Println("dec = ", dec)
	fmt.Println("xxx = ", xxx)
	fmt.Println("div = ", div)
	fmt.Println("__div = ", __div)
}
