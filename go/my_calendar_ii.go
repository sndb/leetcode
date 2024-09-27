package main

type MyCalendarTwo struct {
	intervals1, intervals2 [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start, end int) bool {
	iv := [2]int{start, end}
	for _, iv2 := range c.intervals2 {
		if _, ok := intervalIntersection(iv2, iv); ok {
			return false
		}
	}
	for _, iv1 := range c.intervals1 {
		if iv2, ok := intervalIntersection(iv1, iv); ok {
			c.intervals2 = append(c.intervals2, iv2)
		}
	}
	c.intervals1 = append(c.intervals1, iv)
	return true
}

func intervalIntersection(iv1, iv2 [2]int) (x [2]int, ok bool) {
	if iv2[1] > iv1[0] && iv2[0] < iv1[1] {
		return [2]int{max(iv1[0], iv2[0]), min(iv1[1], iv2[1])}, true
	}
	return [2]int{}, false
}
