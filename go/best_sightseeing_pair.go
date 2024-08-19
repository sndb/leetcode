package main

func maxScoreSightseeingPair(values []int) int {
	curr := 0
	best := 0
	for _, v := range values {
		best = max(best, v+curr)
		curr = max(curr, v) - 1
	}
	return best
}
