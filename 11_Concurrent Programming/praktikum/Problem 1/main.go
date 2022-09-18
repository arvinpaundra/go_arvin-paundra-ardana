package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	// Letter Frequency

	var mutex sync.Mutex

	var words = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	var frequency = make(map[string]int)

	go func(words string, mutex *sync.Mutex) {
		mutex.Lock()

		defer mutex.Unlock()

		// TODO mengubah semua karakter string menjadi lower case
		lowerCase := strings.ToLower(words)

		// TODO memisah semua string menjadi slice of string
		splittedWords := strings.Split(lowerCase, "")

		// TODO loop over slice of string
		for _, word := range splittedWords {
			// TODO memasukkan masing-masing elemen pada slice ke bentuk map
			frequency[word]++
		}
	}(words, &mutex)

	go func(mutex *sync.Mutex) {
		mutex.Lock()

		defer mutex.Unlock()

		// TODO loop over map frequency untuk menghapus karakter selain a-z
		for key := range frequency {
			if key < "a" || key > "z" {
				// TODO hapus key pada map selain a-z
				delete(frequency, key)
			}
		}

		fmt.Println(frequency)
	}(&mutex)

	time.Sleep(1 * time.Second)
}
