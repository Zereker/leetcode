package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 2 {
		return []int{1, 2}
	}

	left, right := 0, len(numbers)-1

	for left < right {
		value := numbers[left] + numbers[right]

		if value == target {
			return []int{left + 1, right + 1}
		} else if value > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

func main() {
	fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 7))
	fmt.Printf("%v", twoSum([]int{5, 25, 75}, 100))
}
