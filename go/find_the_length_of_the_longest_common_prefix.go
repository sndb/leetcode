package main

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	p1 := map[int]bool{}
	for _, x := range arr1 {
		for x > 0 {
			p1[x] = true
			x /= 10
		}
	}
	p2 := map[int]bool{}
	for _, y := range arr2 {
		for y > 0 {
			p2[y] = true
			y /= 10
		}
	}
	l := 0
	for p := range p1 {
		if p2[p] {
			l = max(l, integerLength(p))
		}
	}
	return l
}

func integerLength(x int) int {
	n := 0
	for x > 0 {
		x /= 10
		n++
	}
	return n
}
