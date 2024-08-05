package main

func kthDistinct(arr []string, k int) string {
	h := map[string]int{}
	for _, s := range arr {
		h[s]++
	}
	for _, s := range arr {
		if h[s] == 1 {
			k--
			if k == 0 {
				return s
			}
		}
	}
	return ""
}
