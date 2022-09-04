package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var temp = make(map[string]int)
	var result = make([]pair, 0)

	for _, value := range items {
		temp[value]++
	}

	for key, element := range temp {
		result = append(result, pair{name: key, count: element})
	}

	return result
}

func main() {
	// MOST APPEAR ITEM

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "baksetball", "tenis"}))
}
