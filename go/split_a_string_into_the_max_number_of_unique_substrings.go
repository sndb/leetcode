package main

func maxUniqueSplit(s string) int {
	result := 0
	seen := map[string]bool{}

	var walk func(int, int)
	walk = func(i, n int) {
		if i == len(s) {
			result = max(result, n)
			return
		}
		for j := i + 1; j <= len(s); j++ {
			t := s[i:j]
			if !seen[t] {
				seen[t] = true
				walk(j, n+1)
				seen[t] = false
			}
		}
	}
	walk(0, 0)

	return result
}
