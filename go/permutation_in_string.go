package main

func checkInclusion(s1, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	p1 := [26]int{}
	for _, c := range s1 {
		p1[c-'a']++
	}

	i, j := 0, 0
	p2 := [26]int{}
	for j < len(s1) {
		p2[s2[j]-'a']++
		j++
	}

	for j < len(s2) {
		if p1 == p2 {
			return true
		}
		p2[s2[i]-'a']--
		p2[s2[j]-'a']++
		i++
		j++
	}
	return p1 == p2
}
