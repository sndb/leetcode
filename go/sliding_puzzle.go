package main

func slidingPuzzle(board [][]int) int {
	type grid [2][3]int
	solution := grid{{1, 2, 3}, {4, 5, 0}}

	moves := 0
	queue := []grid{{
		{board[0][0], board[0][1], board[0][2]},
		{board[1][0], board[1][1], board[1][2]},
	}}
	next := []grid{}
	seen := map[grid]bool{}
	for len(queue) > 0 {
		for len(queue) > 0 {
			g := queue[0]
			queue = queue[1:]
			if seen[g] {
				continue
			}
			seen[g] = true

			if g == solution {
				return moves
			}

			var r0, c0 int
			for r := range 2 {
				for c := range 3 {
					if g[r][c] == 0 {
						r0, c0 = r, c
					}
				}
			}

			g1 := g
			g1[0][c0], g1[1][c0] = g1[1][c0], g1[0][c0]
			next = append(next, g1)

			if c0 < 2 {
				g2 := g
				g2[r0][0], g2[r0][1] = g2[r0][1], g2[r0][0]
				next = append(next, g2)
			}

			if c0 > 0 {
				g3 := g
				g3[r0][1], g3[r0][2] = g3[r0][2], g3[r0][1]
				next = append(next, g3)
			}
		}
		queue, next = next, queue
		moves++
	}
	return -1
}
