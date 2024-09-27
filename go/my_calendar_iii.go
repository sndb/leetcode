package main

import "slices"

type MyCalendarThree struct {
	starts []int
	ends   []int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

func (c *MyCalendarThree) Book(start, end int) int {
	c.starts = append(c.starts, start)
	slices.Sort(c.starts)

	c.ends = append(c.ends, end)
	slices.Sort(c.ends)

	i, j, k, maxK := 0, 0, 0, 0
	for i < len(c.starts) {
		if c.starts[i] < c.ends[j] {
			i++
			k++
		} else if c.starts[i] > c.ends[j] {
			j++
			k--
		} else {
			i++
			j++
		}
		maxK = max(maxK, k)
	}
	return maxK
}
