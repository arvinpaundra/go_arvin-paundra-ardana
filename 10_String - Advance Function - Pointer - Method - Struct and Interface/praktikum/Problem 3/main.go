package main

import (
	"fmt"
)

func swap(a, b *int) {
	// TODO simpan value a pada variabel temp
	var temp = *a

	// TODO lakukan swap value a dengan value b, b dengan temp
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
