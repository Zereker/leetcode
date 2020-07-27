package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func nthUglyNumber(n int) int {
	nums := make([]int, 0, 1690)
	var primes = [3]int{2, 3, 5}

	h := new(IntHeap)
	heap.Init(h)
	heap.Push(h, 1)

	seen := make(map[int]struct{}, 0)
	seen[1] = struct{}{}

	for i := 0; i < 1690; i++ {
		currUgly := heap.Pop(h).(int)

		nums = append(nums, currUgly)
		for _, prime := range primes {
			newUgly := currUgly * prime

			if _, ok := seen[newUgly]; !ok {
				seen[newUgly] = struct{}{}
				heap.Push(h, newUgly)
			}
		}
	}

	return nums[n-1]
}

func main() {
	println(nthUglyNumber(11))
}
