package main

import "fmt"

func minimumLength(s string) int {
	h := [26]int{}
	for _, c := range s {
		h[c-'a']++
	}
	r := 0
	for _, n := range h {
		if n > 0 {
			r += 2 - n%2
		}
	}
	return r
}

func main() {
	fmt.Println(minimumLength("abaacbcbb"))
}
