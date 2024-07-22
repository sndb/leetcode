package main

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func maxVowels(s string, k int) int {
	v := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			v++
		}
	}
	w := v
	for i := 0; i+k < len(s); i++ {
		if isVowel(s[i]) {
			v--
		}
		if isVowel(s[i+k]) {
			v++
		}
		w = max(v, w)
	}
	return w
}
