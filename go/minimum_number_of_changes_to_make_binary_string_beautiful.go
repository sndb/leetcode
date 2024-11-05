package main

func minChanges(s string) int {
	c := 0
	for i := 0; i < len(s); i += 2 {
		if p := s[i : i+2]; p == "01" || p == "10" {
			c++
		}
	}
	return c
}
