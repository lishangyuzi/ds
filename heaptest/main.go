package main

import (
	"fmt"
)

type Heap struct {
	source []int
	size   int
}

func adjustHeap(heap Heap, idx int) {
	root := idx
	left := 2*idx + 1
	right := 2*idx + 2
	max := root
	if left < heap.size && heap.source[left] > heap.source[max] {
		max = left
	}
	if right < heap.size && heap.source[right] > heap.source[max] {
		max = right
	}
	if max != root {
		heap.source[root], heap.source[max] = heap.source[max], heap.source[root]
		adjustHeap(heap, max)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	heap := Heap{
		source: arr,
		size:   len(arr),
	}

	// build heap
	for i := heap.size / 2; i >= 0; i-- {
		adjustHeap(heap, i)
	}

	//heap sort
	sorted := []int{}
	length := heap.size
	for i := length; i > 0; i-- {
		fmt.Println("i:", i, "\n")
		sorted = append(sorted, heap.source[0])
		heap.source[0], heap.source[heap.size-1] = heap.source[heap.size-1], heap.source[0]
		heap.source = heap.source[:heap.size-1]
		heap.size--
		adjustHeap(heap, 0)
	}

	fmt.Println(length)
	fmt.Println(sorted)
}
