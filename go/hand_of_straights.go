package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	want := map[int]int{}
	for _, n := range hand {
		if want[n] > 0 {
			want[n]--
		} else {
			for i := 1; i < groupSize; i++ {
				want[n+i]++
			}
			if len(want) > len(hand) {
				return false
			}
		}
	}
	for _, c := range want {
		if c != 0 {
			return false
		}
	}
	return true
}
