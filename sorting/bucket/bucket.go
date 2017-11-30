package bucket

import "container/heap"

type HeapArray []int

func (h HeapArray) Len() int           { return len(h) }
func (h HeapArray) Less(i, j int) bool { return h[i] < h[j] }
func (h HeapArray) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HeapArray) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapArray) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sort(arr []int) {
	bucket := make([]HeapArray, 10)

	max := arr[0]
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}

	min := arr[0]
	for _, elem := range arr {
		if elem < min {
			min = elem
		}
	}

	width := (max-min)/10 + 1

	for i := 0; i < 10; i++ {
		// fmt.Println(i)
		heap.Init(&bucket[i])
	}

	for _, elem := range arr {
		x := (elem - min) / width
		heap.Push(&bucket[x], elem)
	}

	index := 0
	for i := 0; i < 10; i++ {
		m := len(bucket[i])
		for j := 0; j < m; j++ {
			arr[index] = heap.Pop(&bucket[i]).(int)
			index++
		}
	}
}