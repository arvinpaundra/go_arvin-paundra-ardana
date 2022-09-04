package main

import "fmt"

func fibonacci(number int) int {
	if number <= 1 {
		return number
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}
}

func main() {
	// FIBONACCI RECURSIVE

	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
