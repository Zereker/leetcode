package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if len(nums) == 1 {
		if target == nums[0] {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	index := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		value := nums[mid]
		if value == target {
			index = mid
			break
		} else if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if index == -1 {
		return []int{-1, -1}
	}

	i, j := index, index
	for i > 0 {
		value := nums[i-1]
		if value == target {
			i--
		} else {
			break
		}
	}

	for j < len(nums)-1 {
		value := nums[j+1]
		if value == target {
			j++
		} else {
			break
		}
	}

	return []int{i, j}
}

func main() {
	fmt.Printf("%v\n", searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Printf("%v\n", searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Printf("%v\n", searchRange([]int{}, 0))
	fmt.Printf("%v\n", searchRange([]int{2, 2}, 2))
	fmt.Printf("%v\n", searchRange([]int{1, 3}, 1))
}
