package main

import (
	"fmt"
)

func reverse(slice []int) []int {
	var new_slice []int
	for i := len(slice) - 1; i >= 0; i-- {
		new_slice = append(new_slice, slice[i])
	}
	return new_slice
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	var new_slice []int
	reversed_a := reverse(slice1)
	reversed_b := reverse(slice2)

	if len(reversed_a) < len(reversed_b) {
		diff := len(reversed_b) - len(reversed_a)
		for i := 0; i < diff; i++ {
			new_slice = append(new_slice, reversed_b[i])
		}
		reversed_b = reversed_b[diff:]
	} else if len(reversed_a) > len(reversed_b) {
		diff := len(reversed_a) - len(reversed_b)
		for i := 0; i < diff; i++ {
			new_slice = append(new_slice, reversed_a[i])
		}
		reversed_a = reversed_a[diff:]
	}
	for j := 0; j < len(reversed_a); j++ {
		new_slice = append(new_slice, reversed_a[j])
		new_slice = append(new_slice, reversed_b[j])
	}

	return new_slice
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
