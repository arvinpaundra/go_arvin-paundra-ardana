package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var result []int

	for i := 0; i < len(angka); i++ {
		var isExist bool

		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] && i != j {
				isExist = true
				break
			}
		}

		if !isExist {
			str := string(angka[i])
			parse, _ := strconv.Atoi(str)
			result = append(result, parse)
		}
	}

	return result
}

func main() {
	// MUNCUL SEKALI
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
