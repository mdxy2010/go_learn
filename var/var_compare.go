package main

import "fmt"

func main() {

	i, j :=  1, 2
	var m, n int
	m = 6
	n = 9

	if i == 1 || j == 2 {
		fmt.Printf("i = %d j = %d\n", i, j)
		fmt.Printf("i == j\n")
	}

	if i < j {
		fmt.Printf("i < j\n")
	}

	if n > m  {
		fmt.Printf("n > m\n")
	}
}
