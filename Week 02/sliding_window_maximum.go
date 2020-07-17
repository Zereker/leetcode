package main

import "fmt"

func max(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == k {
		return []int{max(nums)}
	}

	var windows, result []int
	for i := 0; i < len(nums); i++ {
		windows = append(windows, nums[i])
		if len(windows) < k {
			continue
		}

		result = append(result, max(windows))
		windows = windows[1:]
	}

	return result
}

func main() {
	result := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	for _, v := range result {
		fmt.Println(v)
	}
}
