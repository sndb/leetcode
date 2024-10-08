package main

func minSwaps(s string) int {
	i1, i2 := 0, len(s)
	ops, cls, sws := 0, 0, 0
	for i1 < i2 {
		if s[i1] == '[' {
			ops++
		} else {
			cls++
		}
		if cls > ops {
			i2--
			for s[i2] != '[' {
				i2--
			}
			cls--
			ops++
			sws++
		}
		i1++
	}
	return sws
}
