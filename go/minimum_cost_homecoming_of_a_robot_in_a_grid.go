package main

import "cmp"

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	cost := 0
	dr := cmp.Compare(homePos[0], startPos[0])
	dc := cmp.Compare(homePos[1], startPos[1])
	for r := startPos[0] + dr; r != homePos[0]+dr; r += dr {
		cost += rowCosts[r]
	}
	for c := startPos[1] + dc; c != homePos[1]+dc; c += dc {
		cost += colCosts[c]
	}
	return cost
}
