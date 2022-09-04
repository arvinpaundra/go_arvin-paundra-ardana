package main

import (
	"fmt"
)

func MaximumBuyProduct(money int, productPrice []int) {
	var total int

	BubbleSort(productPrice)

	for i := 0; i < len(productPrice); i++ {
		money -= productPrice[i]

		if money >= 0 {
			total += 1
		}
	}

	fmt.Println(total)
}

func main() {
	// MAXIMUM BUY PRODUCT

	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	MaximumBuyProduct(0, []int{10000, 30000})
}

func BubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := len(arr) - 1 - 1; j >= 0; j-- {
			if arr[j] <= arr[j+1] {
				continue
			}

			temp := arr[j+1]
			arr[j+1] = arr[j]
			arr[j] = temp
		}
	}

	return arr
}
