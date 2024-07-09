package main

func averageWaitingTime(customers [][]int) float64 {
	total := 0
	free := 0
	for _, c := range customers {
		free = max(free, c[0]) + c[1]
		total += free - c[0]
	}
	return float64(total) / float64(len(customers))
}
