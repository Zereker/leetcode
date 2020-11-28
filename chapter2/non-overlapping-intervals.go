package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var (
		result = 0
		prev   = intervals[0][1]
	)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= prev { // 没有冲突
			prev = intervals[i][1]
		} else { // 若冲突则将区间终点更新为较小的那个，即移除终点值大的区间（终点值大和后续区间冲突可能性也大）
			if intervals[i][1] < prev {
				prev = intervals[i][1]
			}

			result++
		}
	}

	return result
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
}
