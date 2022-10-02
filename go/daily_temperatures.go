package main

func dailyTemperatures(temperatures []int) []int {
	r := make([]int, len(temperatures))
	m := map[int]int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		k, ok := m[temperatures[i]+1]
		if ok {
			r[i] = k - i
		}
		for j := temperatures[i]; j >= 30; j-- {
			m[j] = i
		}
	}
	return r
}
