package main

type MyCalendar struct {
	intervals [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	for _, i := range c.intervals {
		if end > i[0] && start < i[1] {
			return false
		}
	}
	c.intervals = append(c.intervals, [2]int{start, end})
	return true
}
