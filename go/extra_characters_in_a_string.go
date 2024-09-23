package main

func minExtraChar(s string, dictionary []string) int {
	words := map[string]bool{}
	for _, w := range dictionary {
		words[w] = true
	}

	n := len(s)
	m := make([]int, n+1)
	for i := 1; i <= n; i++ {
		m[i] = m[i-1] + 1
		for j := 0; j < i; j++ {
			if w := s[j:i]; words[w] {
				m[i] = min(m[i], m[j])
			}
		}
	}
	return m[n]
}
