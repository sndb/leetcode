package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, c1, c2 := m+n-1, m-1, n-1; i >= 0 && c2 >= 0; i-- {
		if c1 < 0 || nums1[c1] < nums2[c2] {
			nums1[i] = nums2[c2]
			c2--
		} else {
			nums1[i] = nums1[c1]
			c1--
		}
	}
}

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{2, 5, 6}
	merge(n1, 3, n2, 3)
	fmt.Println(n1)
}
