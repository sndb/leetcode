package main

import "strings"

func decodeString(s string) string {
	result := ""
	i := 0
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			n := 0
			for s[i] != '[' {
				n *= 10
				n += int(s[i] - '0')
				i++
			}
			i++
			start := i
			depth := 1
			for depth > 0 {
				if s[i] == '[' {
					depth++
				} else if s[i] == ']' {
					depth--
				}
				i++
			}
			inner := decodeString(s[start : i-1])
			result += strings.Repeat(inner, n)
		} else {
			result += string(s[i])
			i++
		}
	}
	return result
}
