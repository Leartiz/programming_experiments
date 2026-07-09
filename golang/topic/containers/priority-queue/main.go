package main

import (
	"container/heap"
	"fmt"
)

// NOTE: container/heap — min-heap по умолчанию.
// Pop() отдаёт элемент с наименьшим "приоритетом" (Less).
//
// Нужно реализовать heap.Interface:
//   Len, Less, Swap  — на слайсе
//   Push, Pop        — для heap.Push / heap.Pop (Pop без верхнего элемента)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &IntHeap{3, 1, 4, 1, 5}
	heap.Init(h)

	fmt.Println("push 2, pop min:")
	heap.Push(h, 2)
	fmt.Println(heap.Pop(h)) // 1
	fmt.Println(heap.Pop(h)) // 1
	fmt.Println(heap.Pop(h)) // 2

	fmt.Println("drain:")
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
