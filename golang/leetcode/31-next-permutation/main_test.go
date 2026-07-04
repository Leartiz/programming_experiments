package main

import (
	"slices"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "ex1",
			in:   []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "ex2",
			in:   []int{2, 3, 1},
			want: []int{3, 1, 2},
		},
		{
			name: "ex3",
			in:   []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "ex4",
			in:   []int{2, 2, 0, 4, 3, 1},
			want: []int{2, 2, 1, 0, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.in)
			nextPermutation(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("got %v, want %v", nums, tt.want)
			}
		})
	}
}

func TestSlicesReverseSubslice(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		from int
		want []int
	}{
		{name: "tail len 2", in: []int{1, 2, 3, 4}, from: 2, want: []int{1, 2, 4, 3}},
		{name: "tail len 1 - без изменений", in: []int{1, 2, 3}, from: 2, want: []int{1, 2, 3}},
		{name: "весь хвост с from=1", in: []int{1, 2, 3}, from: 1, want: []int{1, 3, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.in)
			slices.Reverse(nums[tt.from:])
			if !slices.Equal(nums, tt.want) {
				t.Errorf("Reverse(nums[%d:]): got %v, want %v", tt.from, nums, tt.want)
			}
		})
	}
}
