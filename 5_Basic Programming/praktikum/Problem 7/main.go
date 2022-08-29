package main

import "fmt"

func playWithAsterisk(n int) {
	for i := 0; i < n; i++ {
		for j := n - i; j > 1; j-- {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}

func main() {
	// PLAY WITH ASTERISK
	playWithAsterisk(5)
}
