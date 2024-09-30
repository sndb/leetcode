package main

type CustomStack struct {
	stack, incrs []int
	len, cap     int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, maxSize),
		incrs: make([]int, maxSize),
		len:   0,
		cap:   maxSize,
	}
}

func (cs *CustomStack) Push(x int) {
	if cs.len == cs.cap {
		return
	}
	cs.stack[cs.len] = x
	cs.len++
}

func (cs *CustomStack) Pop() int {
	if cs.len == 0 {
		return -1
	}
	i := cs.len - 1
	r := cs.stack[i] + cs.incrs[i]
	if i > 0 {
		cs.incrs[i-1] += cs.incrs[i]
	}
	cs.incrs[i] = 0
	cs.len--
	return r
}

func (cs *CustomStack) Increment(k int, val int) {
	if cs.len == 0 {
		return
	}
	cs.incrs[min(k-1, cs.len-1)] += val
}
