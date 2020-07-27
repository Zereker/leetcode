package main

import "container/heap"

type Pair struct {
	Key, Value int
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pairs) Push(x interface{}) {
	*p = append(*p, x.(Pair))
}

func (p *Pairs) Pop() interface{} {
	length := len(*p)

	result := (*p)[length-1]
	*p = (*p)[:length-1]
	return result
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)

	for _, num := range nums {
		m[num] ++
	}

	pairs := new(Pairs)
	heap.Init(pairs)

	for key, value := range m {
		heap.Push(pairs, Pair{
			Key:   key,
			Value: value,
		})
	}

	var result []int
	for i := 0; i < k; i++ {
		pair := heap.Pop(pairs).(Pair)
		result = append(result, pair.Key)
	}

	return result
}
