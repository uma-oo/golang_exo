package main
import (
	"os"
)
const (
	fi3l = 2
)
func validateArgs(args []string) bool {
	if len(args) != 4 {
		return false
	}
	for i := 1; i < 4; i++ {
		if i == fi3l {
			if args[i] != "+" && args[i] != "-" && args[i] != "*" && args[i] != "/" && args[i] != "%" {
				return false
			}
		} else {
			if args[i][0] == '-' {
				for _, c := range args[i][1:] {
					if c < '0' || c > '9' {
						return false
					}
				}
			} else {
				for _, c := range args[i] {
					if c < '0' || c > '9' {
						return false
					}
				}
			}
		}
		if args[i] >= "9223372036854775807" || atoi(args[i]) < -9223372036854775807 {
			return false
		}
	}
	return true
}
func atoi(s string) int {
	var n int
	var sign int = 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n * sign
}
func itoa(n int) string {
	var s string
	g := ""
	if n < 0 {
		g += "-"
		n *= -1
	}
	if n == 0 {
		s += "0"
	}
	for n > 0 {
		s = string(rune(n%10+'0')) + s
		n /= 10
	}
	return g + s
}
func plus(a, b int) int {
	return a + b
}
func nas(a, b int) int {
	return a - b
}
func darb(a, b int) int {
	return a * b
}
func div(a, b int) int {
	if b == 0 {
		return -1
	}
	return a / b
}
func mod(a, b int) int {
	if b == 0 {
		return -1
	}
	return a % b
}
func main() {
	if !validateArgs(os.Args) {
		return
	}
	fi3l := os.Args[fi3l]
	fa3il := atoi(os.Args[1])
	mf3ol_bih := atoi(os.Args[3])
	var result int
	switch fi3l {
	case "+":
		result = plus(fa3il, mf3ol_bih)
	case "-":
		result = nas(fa3il, mf3ol_bih)
	case "*":
		result = darb(fa3il, mf3ol_bih)
	case "/":
		result = div(fa3il, mf3ol_bih)
	case "%":
		result = mod(fa3il, mf3ol_bih)
	}
	file, _ := os.Create("result")
	if result == -1 {
		if fi3l == "/" {
			file.WriteString("No division by 0\n")
		} else if fi3l == "%" {
			file.WriteString("No modulo by 0\n")
		}
	} else {
		file.WriteString(itoa(result) + "\n")
	}
	file, _ = os.Open("result")
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}


//made by saad salah 
