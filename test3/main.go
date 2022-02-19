package main

import "fmt"

func main() {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sub := []int{1, 6, -1, 10}
	fmt.Printf("%v %v %v", arr, sub, IsValidSubsequence(arr, sub))
}

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	return false
}
