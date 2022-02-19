package main

import "fmt"

func main() {
	fmt.Printf("This is cool stuff man %d", bintodec(111))
}

func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fact(n-1)
}

func bintodec(n int) int {
	var dec int = 0
	var pow int = 1
	for n > 0 {
		dec += n % 10 * pow
		pow *= 2
		n /= 10
	}
	return dec
}
