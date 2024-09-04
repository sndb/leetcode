package main

func robotSim(commands []int, obstacles [][]int) int {
	wall := map[[2]int]bool{}
	for _, o := range obstacles {
		wall[[2]int{o[0], o[1]}] = true
	}
	dist := 0
	px, py := 0, 0
	dx, dy := 0, 1
	for _, c := range commands {
		switch c {
		case -2:
			dx, dy = -dy, dx
		case -1:
			dx, dy = dy, -dx
		default:
			for range c {
				nx, ny := px+dx, py+dy
				if wall[[2]int{nx, ny}] {
					break
				}
				px, py = nx, ny
				dist = max(dist, px*px+py*py)
			}
		}
	}
	return dist
}
