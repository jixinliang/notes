package main

import (
	"container/heap"
	"fmt"
)

type heapFloat32 []float32

func (h *heapFloat32) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	newHeap := old[0 : len(old)-1]
	*h = newHeap
	return x
}

func (h *heapFloat32) Push(x interface{}) {
	*h = append(*h, x.(float32))
}

func (h heapFloat32) Len() int {
	return len(h)
}

func (h heapFloat32) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h heapFloat32) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func main() {
	myHeap := &heapFloat32{1.2, 2.1, 3.1, -100.1}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Println("heap size:", size)
	fmt.Println("heap:", *myHeap)

	myHeap.Push(float32(0.2))
	myHeap.Push(float32(0.23))
	fmt.Println("heap size:", size)
	fmt.Println("heap:", *myHeap)
	fmt.Println("len of heap:", len(*myHeap))

	res := myHeap.Pop()
	fmt.Println("Popped value:", res)
	fmt.Println("heap:", *myHeap)
	fmt.Println("len of heap:", len(*myHeap))

}
