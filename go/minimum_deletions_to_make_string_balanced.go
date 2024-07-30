package main

func minimumDeletions(s string) int {
	r := 0
	b := 0
	for _, c := range s {
		if c == 'a' {
			r = min(r+1, b)
		} else {
			b++
		}
	}
	return r
}
