package containers

type BinaryHeap []int

func (b BinaryHeap) Len() int {
	return len(b)
}

func (b BinaryHeap) Less(i, j int) bool {
	return b[i] > b[j]
}

func (b BinaryHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b *BinaryHeap) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *BinaryHeap) Pop() interface{} {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[0 : n-1]
	return x
}
