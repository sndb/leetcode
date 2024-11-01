package main

func makeFancyString(s string) string {
	var r []rune
	var pp, p rune
	for _, c := range s {
		if c != p || c != pp {
			r = append(r, c)
		}
		pp, p = p, c
	}
	return string(r)
}
