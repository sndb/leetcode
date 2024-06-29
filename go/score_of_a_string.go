package main

func scoreOfString(s string) int {
	r := 0
	for i := 1; i < len(s); i++ {
		x := int(s[i]) - int(s[i-1])
		if x < 0 {
			x = -x
		}
		r += x
	}
	return r
}
