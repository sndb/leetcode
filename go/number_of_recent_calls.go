package main

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	n := 0
	for n < len(rc.queue) && rc.queue[n] < t-3000 {
		n++
	}
	rc.queue = append(rc.queue[n:], t)
	return len(rc.queue)
}
