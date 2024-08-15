package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	} else {
		arg := args[0]
		new_str := ""
		for _, char := range arg {
			if (char <= 'm' && char >= 'a') || (char <= 'M' && char >= 'A') {
				new_str += string(char + 13)
			} else if (char <= 'z' && char > 'm') || (char <= 'Z' && char > 'M') {
				new_str += string(char - 13)
			} else {
				new_str += string(char)
			}
		}
		fmt.Println(new_str)
	}
}
