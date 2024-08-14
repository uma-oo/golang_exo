package main

import (
	"fmt"
	"os"
)

func itExists(s string, c string) bool {
	for _, char := range s {
		if string(char) == c {
			return true
		} else {
			continue
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	} else {
		new_string := ""

		for _, char := range args[0] + args[1] {
			if itExists(new_string, string(char)) {
				continue
			} else {
				new_string += string(char)
			}
		}
		fmt.Println(new_string)

	}
}
