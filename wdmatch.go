package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	} else {
		is_printable := false
		arg := args[1]
		for _, char := range args[0] {
			for i, w := range arg {
				if char == w {
					arg = arg[i+1:]
					is_printable = true
					break
				}
				is_printable = false

			}
		}
		if is_printable {
			fmt.Println(args[0])
		}

	}
}
