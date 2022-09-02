package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	var num = pow(x, n/2)

	if n%2 == 0 {
		return num * num
	}

	return num * num * x
}

func main() {
	// FAST EXPONENTIATION

	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
