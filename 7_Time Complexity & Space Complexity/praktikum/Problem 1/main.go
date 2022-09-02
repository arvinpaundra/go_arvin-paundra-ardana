package main

import "fmt"

func primeNumber(number int) bool {

	if number == 1 {

		return false
	} else if number <= 3 {

		return true
	} else if number == 5 {

		return true
	} else {
		for i := number - 1; i > 1; i /= 2 {
			if number%2 == 0 || number%3 == 0 || number%5 == 0 {

				return false
			}
		}
	}

	return true
}

func main() {
	// PRIME NUMBER WITH O(log n) -> LOGARITHMIC

	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(1500450271))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
