package main

import "fmt"

func main() {
	// MENGHITUNG LUAS PERMUKAAN TABUNG

	// input
	var T float64
	var r float64

	const phi float64 = 3.14

	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scanf("%f \n", &T)

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanf("%f \n", &r)

	Lp := 2 * phi * r * (r + T)

	fmt.Println("Luas permukaan tabung: ", Lp)
}
