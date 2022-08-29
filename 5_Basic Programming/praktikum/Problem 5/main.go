package main

import "fmt"

func palindrome(input string) bool {
	var isPalindrome bool
	var reversed string

	for _, value := range input {
		reversed = string(value) + reversed
	}

	if reversed == input {
		isPalindrome = true
	}

	return isPalindrome
}

func main() {
	// PALINDROME
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}
