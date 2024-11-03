package main

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
next:
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[(i+j)%len(s)] != goal[j] {
				continue next
			}
		}
		return true
	}
	return false
}
