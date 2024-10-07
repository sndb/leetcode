package main

func minLength(s string) int {
	stack := []rune{}
	for _, c := range s {
		stack = append(stack, c)
		for len(stack) >= 2 {
			substr := string(stack[len(stack)-2:])
			if substr == "AB" || substr == "CD" {
				stack = stack[:len(stack)-2]
			} else {
				break
			}
		}
	}
	return len(stack)
}
