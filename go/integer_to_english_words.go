package main

import (
	"fmt"
	"strings"
)

func numberToWords(n int) string {
	if n == 0 {
		return "Zero"
	}
	var s []string
	if b := n / 1e9; b != 0 {
		s = append(append(s, convert(b)...), "Billion")
	}
	if m := n / 1e6 % 1e3; m != 0 {
		s = append(append(s, convert(m)...), "Million")
	}
	if t := n / 1e3 % 1e3; t != 0 {
		s = append(append(s, convert(t)...), "Thousand")
	}
	if u := n % 1e3; u != 0 {
		s = append(s, convert(u)...)
	}
	return strings.Join(s, " ")
}

func convert(n int) []string {
	t1 := [...]string{1: "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	t2 := [...]string{2: "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	var r []string
	if m := n / 100; m != 0 {
		r = append(r, t1[m], "Hundred")
	}
	if m := n % 100; m != 0 {
		if m < 20 {
			r = append(r, t1[m])
		} else {
			r = append(r, t2[m/10])
			if k := m % 10; k != 0 {
				r = append(r, t1[k])
			}
		}
	}
	return r
}

func main() {
	fmt.Println(numberToWords(1000666))
}
