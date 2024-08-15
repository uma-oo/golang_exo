package main

import (
	"github.com/01-edu/z01"
)

func Rot14(s string) string {
	new_str := ""
	for _, char := range s {
		if (char <= 'm' && char >= 'a') || (char <= 'M' && char >= 'A') {
			new_str += string(char + 14)
		} else if (char <= 'z' && char > 'm') || (char <= 'Z' && char > 'M') {
			new_str += string(char - 12)
		} else {
			new_str += string(char)
		}
	}
	return new_str
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
