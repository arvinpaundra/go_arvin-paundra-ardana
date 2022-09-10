package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name       string
	NameEncode string
	Score      int
}

type Cipher interface {
	Encode() string
	Decode() string
}

func (s *Student) Encode() string {
	var nameEncode = ""
	var temp []string

	for _, value := range s.Name {
		shift := (int(value)+17-97)%26 + 97
		asciiToString := rune(shift)
		temp = append(temp, string(asciiToString))
	}

	nameEncode = strings.Join(temp, "")

	return nameEncode
}

func (s *Student) Decode() string {
	var nameDecode = ""
	var temp []string

	for _, value := range s.NameEncode {
		shift := (int(value)-17-97)%26 + 97
		fmt.Println(shift)
		asciiToString := rune(shift)
		temp = append(temp, string(asciiToString))
	}

	nameDecode = strings.Join(temp, "")

	return nameDecode
}

func main() {
	var menu int
	var a = Student{}
	var c Cipher = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.Name)
		fmt.Print("\nEncode of Student's Name " + a.Name + " is " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.NameEncode)
		fmt.Print("\nDecode of Student's Name " + a.NameEncode + " is " + c.Decode())
	} else {
		fmt.Print("\nWrong input name menu")
	}
}
