package main

import (
	"fmt"
)

func Slice(a []string, nbrs ...int) []string {
	for i, nbr := range nbrs {
		if nbr < 0 {
			nbr = nbr + len(a)
			nbrs[i] = nbr
		}
	}
	if len(nbrs) == 1 {
		return a[nbrs[0]:]
	} else if len(nbrs) == 2 {
		if nbrs[0] > nbrs[1] {
			return nil
		} else {
			return a[nbrs[0]:nbrs[1]]
		}
	} else {
		return a
	}
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

// []string{"algorithm", "ascii", "package", "golang"}
// []string{"ascii", "package"}
// []string{"ascii", "package", "golang"}
// []string{"package"}
// []string(nil)
