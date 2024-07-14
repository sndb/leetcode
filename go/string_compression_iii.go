package main

import "strings"

func compressedString(word string) string {
	var comp strings.Builder
	for i := 0; i < len(word); i++ {
		char := word[i]
		count := 1
		for i+1 < len(word) && word[i+1] == char && count < 9 {
			count++
			i++
		}
		comp.WriteByte(byte(count + '0'))
		comp.WriteByte(char)
	}
	return comp.String()
}
