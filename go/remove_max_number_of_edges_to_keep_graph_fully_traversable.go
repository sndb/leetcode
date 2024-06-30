package main

import "fmt"

type disjset []int

func makeDisjset(n int) disjset {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	return disjset(a)
}

func (d disjset) clone() disjset {
	c := make(disjset, len(d))
	copy(c, d)
	return c
}

func (d disjset) find(x int) int {
	if d[x] != x {
		d[x] = d.find(d[x])
	}
	return d[x]
}

func (d disjset) union(x, y int) bool {
	rx := d.find(x)
	ry := d.find(y)
	if rx == ry {
		return false
	}
	d[rx] = ry
	return true
}

func (d disjset) ok() bool {
	x := d.find(1)
	for i := 2; i < len(d); i++ {
		if d.find(i) != x {
			return false
		}
	}
	return true
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	total := 0
	ds1 := makeDisjset(n + 1)
	for _, e := range edges {
		if e[0] == 3 && ds1.union(e[1], e[2]) {
			total++
		}
	}
	ds2 := ds1.clone()
	for _, e := range edges {
		if e[0] == 1 && ds1.union(e[1], e[2]) || e[0] == 2 && ds2.union(e[1], e[2]) {
			total++
		}
	}
	if ds1.ok() && ds2.ok() {
		return len(edges) - total
	}
	return -1
}

func main() {
	fmt.Println(maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}))
}
