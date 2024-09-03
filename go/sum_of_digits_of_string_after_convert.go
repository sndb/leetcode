package main

func getLucky(s string, k int) int {
	n := 0
	for _, c := range s {
		n += digitSum(int(c - 'a' + 1))
	}
	for range k - 1 {
		n = digitSum(n)
	}
	return n
}

func digitSum(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}
