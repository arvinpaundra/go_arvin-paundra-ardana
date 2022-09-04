package main

import "fmt"

func maxSequence(arr []int) int {
	const (
		UintSize = 32 << (^uint(0) >> 32 & 1)
		MaxInt   = 1<<(UintSize-1) - 1
		MinInt   = -MaxInt - 1
	)

	var currentMax = 0
	var max = MinInt

	for i := 0; i < len(arr)-1; i++ {
		currentMax += arr[i]

		if currentMax > max {
			max = currentMax
		}

		if currentMax < 0 {
			currentMax = 0
		}
	}

	return max
}

func main() {
	// TOTAL MAKSIMUM DARI DERET BILANGAN

	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(maxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(maxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
