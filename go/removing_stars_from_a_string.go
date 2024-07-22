package main

func removeStars(s string) string {
	q := []rune{}
	for _, c := range s {
		if c == '*' {
			q = q[:len(q)-1]
		} else {
			q = append(q, c)
		}
	}
	return string(q)
}
