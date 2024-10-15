package main

func minimumSteps(s string) int64 {
	var r, b int64
	for _, c := range s {
		if c == '1' {
			b++
		} else {
			r += b
		}
	}
	return r
}
