package main

import "fmt"

func fractionAddition(expr string) string {
	n1 := 0
	d1 := 1
	s := 1
	i := 0
	for i < len(expr) {
		switch expr[i] {
		case '+':
			s = 1
			i++
		case '-':
			s = -1
			i++
		default:
			a := i
			for i < len(expr) && expr[i] != '+' && expr[i] != '-' {
				i++
			}
			var n2, d2 int
			fmt.Sscanf(expr[a:i], "%d/%d", &n2, &d2)
			n1 = n1*d2 + s*n2*d1
			d1 = d1 * d2
		}
	}
	g := gcd(n1, d1)
	return fmt.Sprintf("%d/%d", n1/g, d1/g)
}

func gcd(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(fractionAddition("-1/2+1/2"))
	fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("1/3-1/2"))
}
