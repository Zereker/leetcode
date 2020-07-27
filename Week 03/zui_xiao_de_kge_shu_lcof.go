package main

import (
	"container/heap"
	"fmt"
)

func getLeastNumbers(arr []int, k int) []int {
	length := len(arr)
	if arr == nil || k > length || length <= 0 || k <= 0 {
		return []int{}
	}

	h := &IntHeap{}
	heap.Init(h)

	for _, v := range arr {
		heap.Push(h, v)
	}

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}

	return res
}

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

func (h *IntHeap) Max() int {
	return (*h)[0]
}

func main() {
	ints := getLeastNumbers([]int{0, 1, 2, 1}, 4)
	fmt.Printf("%v", ints)
}
