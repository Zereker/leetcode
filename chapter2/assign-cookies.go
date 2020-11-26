package main

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 || len(g) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)

	sIndex := len(s) - 1
	gIndex := len(g) - 1

	var result int
	for sIndex >= 0 && gIndex >= 0 {
		if s[sIndex] >= g[gIndex] {
			result++
			sIndex--
			gIndex--
		} else {
			gIndex--
		}
	}

	return result
}

func main() {
	println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	println(findContentChildren([]int{1, 2, 3}, []int{3}))
}
