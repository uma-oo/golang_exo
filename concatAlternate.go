package main

import (
	"fmt"
)

func ConcatAlternate(slice1, slice2 []int) []int {
	var new_slice []int
	if len(slice2) > len(slice1) {
		slice1, slice2 = slice2, slice1
	}
	for i,v :=range slice1{
		new_slice=append(new_slice, v)
		if i<len(slice2){
			new_slice=append(new_slice, slice2[i])
		}
	}

	return new_slice
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
