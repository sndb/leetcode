package main

import (
	"fmt"
	"slices"
)

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	const mod = 1000000007
	const base = 26
	pow := 1
	h1 := 0
	h2 := 0
	n := 0
	for i := 0; i < len(s); i++ {
		h1 = (h1*base + int(s[i]-'a'+1)) % mod
		h2 = (h2 + int(s[i]-'a'+1)*pow) % mod
		pow = (pow * base) % mod
		if h1 == h2 {
			n = i
		}
	}

	r := []byte(s[n+1:])
	slices.Reverse(r)
	return string(r) + s
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}
