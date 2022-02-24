package main

import "fmt"

type A struct {
	val1 int
	val2 string
}

func main() {
	//var a A = A{10, "hello"}

	data := []int{10, 20, 30}
	simple(data)
	fmt.Printf("%v ", data)
}

func simple(a []int) {
	a[0] = 20
	a = append(a, 40)
	fmt.Printf("%v", a)
}
