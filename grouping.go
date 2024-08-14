package main

import (
	"fmt"
	"os"
)

func split(s string) []string {
	word := ""
	var slice []string
	for i, char := range s {
		if s[i] != ' ' && s[i] != ',' && s[i] != '\t' {
			word += string(char)
		} else if word != "" {
			slice = append(slice, word)
			word = ""
		}
	}
	if word != "" {
		slice = append(slice, word)
	}
	return slice
}

func handle_regex(s string) []string {
	word := ""
	var slices []string
	for _, char := range s {
		if char != '|' {
			word += string(char)
		} else if word != "" {
			slices = append(slices, word)
			word = ""
		}
	}
	if word != "" {
		slices = append(slices, word)
	}

	return slices
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	} else {
		m := make(map[string][]string)
		slice := split(args[1])
		if args[0][0] != '(' || args[0][len(args[0])-1] != ')' {
			return
		} else {
			to_find := handle_regex(args[0][1 : len(args[0])-1])

			for i := 0; i < len(to_find); i++ {
				for j := 0; j < len(slice); j++ {
					for k := 0; k < len(slice[j]); k++ {
						if len(to_find[i])+k <= len(slice[j]) && string(to_find[i]) == string(slice[j][k:len(to_find[i])+k]) {
							if !contains(m[string(to_find[i])], slice[j]) {
								if _, ok := m[string(to_find[i])]; !ok {
									m[string(to_find[i])] = []string{}
								}
								m[string(to_find[i])] = append(m[string(to_find[i])], slice[j])
							}
							break // Break the innermost loop, but continue checking other words
						}
					}
				}
			}
			// let's print things in order !!!
			i := 0
			for _, word := range slice {
				for _, value := range m {
					for _, w := range value {
						if word == w {
							fmt.Printf("%d: %s\n", i+1, word)
							i++
						}
					}
				}
			}
		}
	}
}
