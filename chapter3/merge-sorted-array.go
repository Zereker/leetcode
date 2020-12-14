package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	first := nums1[0:m]
	second := nums2[0:n]
	nums1 = append(first, second...)

	sort.Ints(nums1)
}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
}
