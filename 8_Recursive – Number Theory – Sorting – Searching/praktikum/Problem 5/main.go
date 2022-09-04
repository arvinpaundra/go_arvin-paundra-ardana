package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FindMinAndMax(arr []int) string {
	var min = arr[0]
	var max = arr[0]
	var minIdx int
	var maxIdx int

	for i, value := range arr {
		if value > max {
			max = value
			maxIdx = i
		}

		if value < min {
			min = value
			minIdx = i
		}
	}

	var result = strings.Join([]string{"min:", strconv.Itoa(min), "index:", strconv.Itoa(minIdx), "max:", strconv.Itoa(max), "index:", strconv.Itoa(maxIdx)}, " ")

	return result
}

func main() {
	// FIND MIN AND MAX NUMBER

	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}
