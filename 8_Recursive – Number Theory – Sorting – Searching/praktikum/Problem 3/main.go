package main

import "fmt"

func primaSegiEmpat(high, wide, start int) {
	var total = high * wide
	var primes = make([]int, 0)
	var count int
	var index = start + 1
	var sum int

	for index > 1 {
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
				if count == total {
					break
				}
			}
		}

		index++
	}

	for i := 0; i <= len(primes)-1; i++ {
		fmt.Printf("%d ", primes[i])

		if (i+1)%high == 0 {
			fmt.Println()
		}

		sum += primes[i]
	}

	fmt.Println(sum)
}

func main() {
	// PRIMA SEGI EMPAT

	// primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}
