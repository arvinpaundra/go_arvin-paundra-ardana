package main

import "fmt"

func pangkat(base int, pangkat int) int {
	result := base

	for i := 0; i < pangkat-1; i++ {
		result *= base
	}

	return result
}

func main() {
	// EXPONENTIATION
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}
