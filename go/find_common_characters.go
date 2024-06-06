package main

func commonChars(words []string) []string {
	a := [26]int{}
	for i, w := range words {
		h := [26]int{}
		for _, r := range w {
			k := r - 'a'
			if a[k] > 0 || i == 0 {
				h[k]++
				a[k]--
			}
		}
		a = h
	}
	r := []string{}
	for i, c := range a {
		for range c {
			r = append(r, string(rune('a'+i)))
		}
	}
	return r
}
