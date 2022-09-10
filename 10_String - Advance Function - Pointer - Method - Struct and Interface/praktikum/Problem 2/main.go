package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	// TODO menampung input yang telah diubah
	var substituted []string

	// TODO lakukan loop over input
	for _, value := range input {
		// TODO geser karakter sesuai dengan offset
		shift := (int(value)+offset-97)%26 + 97
		// TODO ubah karakter yang telah digeser menjadi string
		asciiToString := rune(shift)
		// TODO gabungkan setiap karakter ke dalam slice
		substituted = append(substituted, string(asciiToString))
	}

	// TODO ubah bentuk slice of string ke bentuk string
	input = strings.Join(substituted, "")

	return input
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
