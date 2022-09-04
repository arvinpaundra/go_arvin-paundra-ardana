package main

import "fmt"

func pairSum(arr []int, target int) []int {
	var arrTemp []int

	for i := range arr {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				arrTemp = append(arrTemp, i, j)
				return arrTemp
			}
		}
	}

	return arrTemp
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))
}
