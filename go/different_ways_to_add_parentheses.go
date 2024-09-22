package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expr string) []int {
	n, err := strconv.Atoi(expr)
	if err == nil {
		return []int{n}
	}

	var results []int
	for i, c := range expr {
		if c != '+' && c != '-' && c != '*' {
			continue
		}
		lvals := diffWaysToCompute(expr[:i])
		rvals := diffWaysToCompute(expr[i+1:])
		for _, v1 := range lvals {
			for _, v2 := range rvals {
				var r int
				switch expr[i] {
				case '+':
					r = v1 + v2
				case '-':
					r = v1 - v2
				case '*':
					r = v1 * v2
				}
				results = append(results, r)
			}
		}
	}
	return results
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
