package main

import "fmt"

// Fact : calculates factorial
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func runFact() {
	fmt.Println(fact(4))
}
