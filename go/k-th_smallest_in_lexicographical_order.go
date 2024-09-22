package main

func findKthNumber(n int, k int) int {
	x := 1
	k -= 1
	for k > 0 {
		next := x + 1
		steps := countSteps(n, x, next)
		if steps <= k {
			x = next
			k -= steps
		} else {
			x *= 10
			k -= 1
		}
	}
	return x
}

func countSteps(n, p1, p2 int) int {
	steps := 0
	for p1 <= n {
		steps += min(n+1, p2) - p1
		p1 *= 10
		p2 *= 10
	}
	return steps
}
