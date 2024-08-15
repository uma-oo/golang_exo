package piscine
import "github.com/01-edu/z01"
func checkBase(base string) bool {
	if len(base) >= 2 && len(base) <= 20 {
		result := true
		for i := 0; i < len(base); i++ {
			for j := i + 1; j < len(base); j++ {
				if base[i] == base[j] || base[i] == '+' || base[i] == '-' {
					result = false
					break
				} else {
					continue
				}
			}
		}
		return result
	} else {
		return false
	}
}
func PrintNbrBase(nbr int, base string) {
	if checkBase(base) {
		var r int
		var s []int
		if nbr < 0 {
			nbr = nbr * -1
			z01.PrintRune('-')
		}
		for nbr > 0 {
			r = nbr % len(base)
			s = append(s, r)
			nbr = nbr / len(base)
		}
		if nbr == -9223372036854775808 {
			PrintNbrBase(922337203685477580, base)
			z01.PrintRune('8')
			return
		}
		for i := len(s) - 1; i >= 0; i-- {
			z01.PrintRune(rune(base[s[i]]))
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}

