package main

func lengthOfLongestSubstring(s string) int {
	seen := map[rune]int{}
	start, length := 0, 0
	for i, c := range s {
		k, ok := seen[c]
		if ok && start <= k {
			start = k + 1
		} else {
			x := i - start + 1
			if x > length {
				length = x
			}
		}
		seen[c] = i
	}
	return length
}
