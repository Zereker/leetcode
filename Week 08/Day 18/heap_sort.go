package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type HeapInt struct {
	sort.IntSlice
}

func (a *HeapInt) Push(x interface{}) {
	a.IntSlice = append(a.IntSlice, x.(int))
}

func (a *HeapInt) Pop() interface{} {
	length := a.IntSlice.Len()
	result := a.IntSlice[length-1]

	a.IntSlice = a.IntSlice[0 : length-1]
	return result
}

func heapSort(arr []int) []int {
	slice := sort.IntSlice(arr)

	var heapInt HeapInt
	heapInt.IntSlice = slice

	heap.Init(&heapInt)
	var result []int
	for heapInt.Len() != 0 {
		data := heap.Pop(&heapInt).(int)
		result = append(result, data)
	}

	return result
}

func main() {
	fmt.Printf("%v", heapSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
