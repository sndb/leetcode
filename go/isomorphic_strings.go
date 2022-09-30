package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	st, ts := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < len(s); i++ {
		a, k := ts[t[i]]
		b, l := st[s[i]]
		if k != l || k && l && (a != s[i] || b != t[i]) {
			return false
		}
		ts[t[i]], st[s[i]] = s[i], t[i]
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))   // true
	fmt.Println(isIsomorphic("badc", "baba")) // false
}
