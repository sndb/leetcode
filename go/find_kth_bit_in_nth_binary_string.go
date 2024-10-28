package main

func findKthBit(n int, k int) byte {
	s := "0"
	for n > 1 {
		s += "1" + reverseInvert(s)
		n--
	}
	return s[k-1]
}

func reverseInvert(s string) string {
	n := len(s)
	t := make([]byte, n)
	for i := 0; i < n; i++ {
		t[i] = '1' - s[n-i-1] + '0'
	}
	return string(t)
}
