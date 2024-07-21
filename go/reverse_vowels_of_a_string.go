package main

func isVowel(c byte) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	l := []byte(s)
	i, j := 0, len(l)-1
	for i < j {
		for i < j && !isVowel(l[i]) {
			i++
		}
		for i < j && !isVowel(l[j]) {
			j--
		}
		l[i], l[j] = l[j], l[i]
		i++
		j--
	}
	return string(l)
}
