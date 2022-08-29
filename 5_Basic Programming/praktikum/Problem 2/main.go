package main

import "fmt"

func main() {
	// GRADE NILAI

	// input
	var studentScore int
	var studentName string
	var result string

	fmt.Print("Masukkan Nama Siswa: ")
	fmt.Scanf("%s \n", &studentName)

	fmt.Print("Masukkan Nilai Siswa: ")
	fmt.Scanf("%d \n\n", &studentScore)

	fmt.Println(studentName)

	if studentScore >= 80 && studentScore <= 100 {
		result = "A"
		fmt.Printf("Nilai %s", result)
	} else if studentScore >= 65 && studentScore <= 79 {
		result = "B"
		fmt.Printf("Nilai %s", result)
	} else if studentScore >= 50 && studentScore <= 64 {
		result = "C"
		fmt.Printf("Nilai %s", result)
	} else if studentScore >= 35 && studentScore <= 49 {
		result = "D"
		fmt.Printf("Nilai %s", result)
	} else if studentScore >= 0 && studentScore <= 34 {
		result = "E"
		fmt.Printf("Nilai %s", result)
	} else {
		fmt.Println("Nilai Invalid")
	}
}
