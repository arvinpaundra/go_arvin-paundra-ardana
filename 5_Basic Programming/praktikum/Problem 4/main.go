package main

import "fmt"

func primeNumber(number int) bool {
	var isPrime bool

	if number == 1 {
		isPrime = false
	} else if number == 2 {
		isPrime = true
	} else {
		for i := 2; i < number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			} else {
				isPrime = true
			}
		}
	}

	return isPrime
}

func main() {
	// BILANGAN PRIMA
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
