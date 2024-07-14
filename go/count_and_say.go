package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		char := s[i]
		count := 1
		for i+1 < len(s) && s[i+1] == char {
			count++
			i++
		}
		b.WriteString(strconv.Itoa(count))
		b.WriteByte(char)
	}
	return b.String()
}

func main() {
	fmt.Println(countAndSay(30))
}
