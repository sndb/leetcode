package main

func maximumGain(s string, x int, y int) int {
	r := 0
	f := func(t string, v int) {
		a := []byte{}
		for _, c := range s {
			a = append(a, byte(c))
			n := len(a)
			if n >= 2 && a[n-2] == t[0] && a[n-1] == t[1] {
				a = a[:n-2]
				r += v
			}
		}
		s = string(a)
	}
	if x > y {
		f("ab", x)
		f("ba", y)
	} else {
		f("ba", y)
		f("ab", x)
	}
	return r
}
