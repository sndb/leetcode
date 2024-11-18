package main

func decrypt(code []int, k int) []int {
	r := make([]int, len(code))
	for i := range code {
		v := 0
		if k > 0 {
			for j := range k {
				v += code[(i+j+1)%len(code)]
			}
		} else {
			for j := range -k {
				v += code[(i-j-1+len(code))%len(code)]
			}
		}
		r[i] = v
	}
	return r
}
