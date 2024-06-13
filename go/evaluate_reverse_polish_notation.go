package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, tok := range tokens {
		n := len(stack)
		switch tok {
		case "+":
			stack[n-2] += stack[n-1]
			stack = stack[:n-1]
		case "-":
			stack[n-2] -= stack[n-1]
			stack = stack[:n-1]
		case "*":
			stack[n-2] *= stack[n-1]
			stack = stack[:n-1]
		case "/":
			stack[n-2] /= stack[n-1]
			stack = stack[:n-1]
		default:
			v, _ := strconv.Atoi(tok)
			stack = append(stack, v)
		}
	}
	return stack[0]
}
