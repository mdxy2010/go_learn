package main

import "fmt"

func main() {
	const a  = 5;
	const b, c , d = 1, 2, 3;
	const u, v  float64 = 0,3;
	
	const (
		d0 = iota
		d1 = iota
		d2 = iota
	)

	const (
		b0 = iota
		b1
		b2
	)
	fmt.Println(u, v, a, b, c, d);
	fmt.Println(d0, d1, d2);
	fmt.Println(b0, b1, b2);
}
