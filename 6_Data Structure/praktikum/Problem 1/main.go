package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var temp []string

	for _, valB := range arrayB {
		var isExist bool

		for _, valA := range arrayA {
			if valA == valB {
				isExist = true
				break
			} else {
				isExist = false
			}
		}

		if !isExist {
			temp = append(temp, valB)
		}
	}

	arrayA = append(arrayA, temp...)

	return arrayA
}

func main() {
	// ARRAY MERGE
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
