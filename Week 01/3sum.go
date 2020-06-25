package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	length := len(nums)
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}

		sum := 0 - nums[i]
		left, right := i+1, length-1
		for left != right {
			if nums[left]+nums[right] < sum {
				left++
			} else if nums[left]+nums[right] > sum {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
			}
		}
	}

	return duplicateRemoval(result)
}

type node struct {
	x, y, z int
}

func duplicateRemoval(data [][]int) [][]int {
	m := make(map[node]struct{}, len(data))

	for _, value := range data {
		d := node{
			x: value[0],
			y: value[1],
			z: value[2],
		}

		if _, ok := m[d]; ok {
			continue
		}

		m[d] = struct{}{}
	}

	var result [][]int
	for key, _ := range m {
		result = append(result, []int{key.x, key.y, key.z})
	}

	return result
}

func main() {
	sum := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("%v\n", sum)
}
