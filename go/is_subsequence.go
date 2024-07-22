package main

func isSubsequence(s string, t string) bool {
	k := 0
	for i := 0; i < len(t) && k < len(s); i++ {
		if s[k] == t[i] {
			k++
		}
	}
	return k == len(s)
}
