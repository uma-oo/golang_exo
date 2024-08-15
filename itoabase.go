package main

import "fmt"

func main() {
	fmt.Println(ItoaBase(15, 16))
	print("klf")
}

func CreateBase(n int) string {
	base := ""
	max_base := "0123456789ABCDEF"
	for i := 0; i < n; i++ {
		base += string(max_base[i])
	}
	return base
}


func ItoaBase(value, base int) string {
	nbr := ""
	for value > 0 {
		nbr = string(rune(value%base+'0')) + nbr
		value = value / base
	}
	return nbr
}

