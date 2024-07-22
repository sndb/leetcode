package main

func isDivisor(d, s string) bool {
	if len(s)%len(d) != 0 {
		return false
	}
	for i := 0; i < len(s); i += len(d) {
		if s[i:i+len(d)] != d {
			return false
		}
	}
	return true
}

func gcdOfStrings(s1 string, s2 string) string {
	for n := min(len(s1), len(s2)); n > 0; n-- {
		if d := s1[:n]; isDivisor(d, s1) && isDivisor(d, s2) {
			return d
		}
	}
	return ""
}
