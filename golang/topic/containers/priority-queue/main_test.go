package main

import (
	"container/heap"
	"slices"
	"testing"
)

func popAll(h *IntHeap) []int {
	var out []int
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(int))
	}
	return out
}

func TestIntHeapPopOrder(t *testing.T) {
	tests := []struct {
		name string
		init []int
		push []int
		want []int
	}{
		{
			name: "min first",
			init: []int{3, 1, 4, 1, 5},
			want: []int{1, 1, 3, 4, 5},
		},
		{
			name: "with extra push",
			init: []int{3, 1, 4, 1, 5},
			push: []int{2},
			want: []int{1, 1, 2, 3, 4, 5},
		},
		{
			name: "single",
			init: []int{42},
			want: []int{42},
		},
		{
			name: "already sorted",
			init: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := IntHeap(slices.Clone(tt.init))
			hp := &h
			heap.Init(hp)

			for _, x := range tt.push {
				heap.Push(hp, x)
			}

			got := popAll(hp)
			if !slices.Equal(got, tt.want) {
				t.Errorf("pop order: got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeapMinAtRoot(t *testing.T) {
	h := &IntHeap{5, 3, 7, 1}
	heap.Init(h)

	if (*h)[0] != 1 {
		t.Errorf("root after Init: got %d, want 1", (*h)[0])
	}
}
