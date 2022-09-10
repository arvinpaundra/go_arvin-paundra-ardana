package main

import (
	"fmt"
	"strings"
)

func ContainSubstr(str, substr string) string {
	// TODO cek substring pada string, jika ada return substring
	if exist := strings.Contains(str, substr); exist {
		return substr
	}

	return "Not match"
}

func Compare(a, b string) string {
	// TODO cek mana string yang lebih panjang antara a atau b
	// TODO jika string b lebih panjang, return fungsi ContainSubstr dengan parameter b diawal
	if len(a) < len(b) {
		return ContainSubstr(b, a)
	} else {
		// TODO jika string a lebih panjang, return fungsi ContainSubstr dengan parameter a diawal
		return ContainSubstr(a, b)
	}
}

func main() {
	// COMPARE STRING

	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
