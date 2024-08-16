package main

func maxDistance(arrays [][]int) int {
	dist := 0
	minV := arrays[0][0]
	maxV := arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		minI := arrays[i][0]
		maxI := arrays[i][len(arrays[i])-1]
		dist = max(dist, maxI-minV, maxV-minI)
		minV = min(minV, minI)
		maxV = max(maxV, maxI)
	}
	return dist
}
