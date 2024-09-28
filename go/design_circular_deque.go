package main

type MyCircularDeque struct {
	data  []int
	index int
	size  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{data: make([]int, k)}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	d.size++
	d.data[d.index] = value
	d.index = (d.index - 1 + len(d.data)) % len(d.data)
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	d.size++
	d.data[(d.index+d.size)%len(d.data)] = value
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.index = (d.index + 1) % len(d.data)
	d.size--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.size--
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[(d.index+1)%len(d.data)]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[(d.index+d.size)%len(d.data)]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == len(d.data)
}
