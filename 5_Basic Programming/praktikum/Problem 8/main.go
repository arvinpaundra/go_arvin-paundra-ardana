package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d\t", i*j)
		}

		fmt.Println()
	}
}

func main() {
	// CETAK TABEL PERKALIAN
	cetakTablePerkalian(9)
}
