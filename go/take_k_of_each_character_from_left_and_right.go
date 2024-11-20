package main

func takeCharacters(s string, k int) int {
	chars := [3]int{}
	for _, c := range s {
		chars[c-'a']++
	}
	for _, n := range chars {
		if n < k {
			return -1
		}
	}
	result := len(s)
	for l, r := 0, len(s); l < len(s); l++ {
		for r > 0 && chars[s[len(s)-r]-'a'] > k {
			chars[s[len(s)-r]-'a']--
			r--
		}
		result = min(result, l+r)
		chars[s[l]-'a']++
	}
	return result
}
