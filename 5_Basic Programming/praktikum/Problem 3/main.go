package main

import "fmt"

func main() {
	// FAKTOR BILANGAN
	var bilangan int

	fmt.Scanf("%d \n\n", &bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i != 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
