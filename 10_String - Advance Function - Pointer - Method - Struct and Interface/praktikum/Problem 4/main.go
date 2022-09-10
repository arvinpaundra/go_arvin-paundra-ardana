package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	// TODO simpan nilai awal min dan max dengan bilangan pertama slice numbers
	min = *numbers[0]
	max = *numbers[0]

	// TODO lakukan loop over numbers
	for _, value := range numbers {
		// TODO cek bila value lebih dari max, maka ganti nilai max sebelumnya dengan value
		if *value > max {
			max = *value
		}

		// TODO cek bila value kurang dari min, maka ganti nilai min sebelumnya dengan value
		if *value < min {
			min = *value
		}
	}

	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Printf("\n%d is the maximum number\n", max)
	fmt.Printf("%d is the minimum number", min)
}
