package main

func sumOfSquares(n int) int {
	r := 0
	for n > 0 {
		x := n % 10
		r += x * x
		n /= 10
	}
	return r
}

func isHappy(n int) bool {
	seen := map[int]struct{}{n: {}}
	for n != 1 {
		n = sumOfSquares(n)
		if _, ok := seen[n]; ok {
			return false
		}
		seen[n] = struct{}{}
	}
	return true
}
