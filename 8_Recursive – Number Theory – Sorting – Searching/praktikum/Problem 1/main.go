package main

import "fmt"

func primeX(number int) int {
	var primes = make([]int, 0)
	var count int
	var index = 2

	if number == 1 {
		return 2
	} else {
		for index > 0 {
			if index == 2 || index == 3 {
				primes = append(primes, index)
				count++
			} else if index == 5 {
				primes = append(primes, index)
				count++
			} else {
				if index%2 != 0 && index%3 != 0 && index%5 != 0 {

					primes = append(primes, index)
					count++
					if count == number {
						break
					}
				}
			}

			index++
		}
	}

	return primes[number-1]
}

func main() {
	// PRIMA KE X

	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
