package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	p1 := m - 1
	p2 := n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}

		p--
	}

	// 只需要考虑 num2 没有插完的即可。因为 num1 没有插完，就不需要动呀
	for p2 >= 0 {
		nums1[p2] = nums2[p2]
		p2--
	}
}

func main() {
	num1 := []int{0}
	num2 := []int{1}
	merge(num1, 0, num2, 1)

	for _, v := range num1 {
		println(v)
	}
}
