package main

var prime = func() []bool {
	t := make([]bool, 1000)
	for p := 2; p < len(t); p++ {
		t[p] = true
	}
	for p := 2; p < len(t); p++ {
		if t[p] {
			for q := p + p; q < len(t); q += p {
				t[q] = false
			}
		}
	}
	return t
}()

func primeSubOperation(nums []int) bool {
	for p := nums[0] - 1; p >= 2; p-- {
		if prime[p] {
			nums[0] -= p
			break
		}
	}
	for i := 1; i < len(nums); i++ {
		for p := nums[i] - nums[i-1] - 1; p >= 2; p-- {
			if prime[p] {
				nums[i] -= p
				break
			}
		}
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	return true
}
