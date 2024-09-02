package main

func chalkReplacer(chalk []int, k int) int {
	s := 0
	for _, c := range chalk {
		s += c
	}
	k %= s
	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return 0
}
