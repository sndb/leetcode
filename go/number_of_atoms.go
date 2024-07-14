package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseFormula(s string) map[string]int {
	result := map[string]int{}
	pending := map[string]int{}
	for i := 0; ; i++ {
		if pending != nil {
			mult := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				mult *= 10
				mult += int(s[i] - '0')
				i++
			}
			for e, n := range pending {
				result[e] += n * max(mult, 1)
			}
			clear(pending)
		}
		if i == len(s) {
			break
		}
		if s[i] == '(' {
			start := i
			depth := 1
			for depth > 0 {
				i++
				if s[i] == '(' {
					depth++
				} else if s[i] == ')' {
					depth--
				}
			}
			pending = parseFormula(s[start+1 : i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			start := i
			i++
			for i < len(s) && s[i] >= 'a' && s[i] <= 'z' {
				i++
			}
			name := s[start:i]
			i--
			pending[name] = 1
		}
	}
	return result
}

func countOfAtoms(formula string) string {
	l := []string{}
	for k, n := range parseFormula(formula) {
		c := ""
		if n > 1 {
			c = strconv.Itoa(n)
		}
		l = append(l, k+c)
	}
	sort.Strings(l)
	return strings.Join(l, "")
}

func main() {
	fmt.Println(countOfAtoms("H2O"))
	fmt.Println(countOfAtoms("Mg(OH)2"))
	fmt.Println(countOfAtoms("K4(ON(SO3)2)2"))
	fmt.Println(countOfAtoms("Mg(OH)2"))
}
