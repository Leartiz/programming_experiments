package main_test

import (
	"reflect"
	"testing"

	flattenconstructor "golang/leetcode/341-flatten-nested-list-iterator/flatten-constructor"
	lazypreload "golang/leetcode/341-flatten-nested-list-iterator/lazy-preload"
	"golang/leetcode/341-flatten-nested-list-iterator/nestedinteger"
	recursivereslice "golang/leetcode/341-flatten-nested-list-iterator/recursive-reslice"
	"golang/leetcode/341-flatten-nested-list-iterator/stack"
)

type iterator interface {
	Next() int
	HasNext() bool
}

type implCase struct {
	name    string
	skip    bool
	newIter func([]*nestedinteger.NestedInteger) iterator
}

func collect(it iterator) []int {
	var got []int
	for it.HasNext() {
		got = append(got, it.Next())
	}
	return got
}

func TestNestedIterator(t *testing.T) {
	tests := []struct {
		name string
		list []*nestedinteger.NestedInteger
		want []int
	}{
		{
			name: "[[1,1],2,[1,1]]",
			list: []*nestedinteger.NestedInteger{
				nestedinteger.NewList([]*nestedinteger.NestedInteger{nestedinteger.NewInteger(1), nestedinteger.NewInteger(1)}),
				nestedinteger.NewInteger(2),
				nestedinteger.NewList([]*nestedinteger.NestedInteger{nestedinteger.NewInteger(1), nestedinteger.NewInteger(1)}),
			},
			want: []int{1, 1, 2, 1, 1},
		},
		{
			name: "empty root",
			list: []*nestedinteger.NestedInteger{},
			want: nil,
		},
		{
			name: "[[]]",
			list: []*nestedinteger.NestedInteger{
				nestedinteger.NewList([]*nestedinteger.NestedInteger{
					nestedinteger.NewEmptyList(),
				}),
			},
			want: nil,
		},
		{
			name: "[[[]]]",
			list: []*nestedinteger.NestedInteger{
				nestedinteger.NewList([]*nestedinteger.NestedInteger{
					nestedinteger.NewList([]*nestedinteger.NestedInteger{
						nestedinteger.NewEmptyList(),
					}),
				}),
			},
			want: nil,
		},
		{
			name: "[1]",
			list: []*nestedinteger.NestedInteger{
				nestedinteger.NewInteger(1),
			},
			want: []int{1},
		},
		{
			name: "[1,[4,[6]]]",
			list: []*nestedinteger.NestedInteger{
				nestedinteger.NewInteger(1),
				nestedinteger.NewList([]*nestedinteger.NestedInteger{
					nestedinteger.NewInteger(4),
					nestedinteger.NewList([]*nestedinteger.NestedInteger{nestedinteger.NewInteger(6)}),
				}),
			},
			want: []int{1, 4, 6},
		},
		{
			name: "empty list between integers",
			list: []*nestedinteger.NestedInteger{
				nestedinteger.NewInteger(2),
				nestedinteger.NewEmptyList(),
				nestedinteger.NewInteger(1),
			},
			want: []int{2, 1},
		},
	}

	impls := []implCase{
		{
			name: "recursive-reslice",
			skip: true,
			newIter: func(list []*nestedinteger.NestedInteger) iterator {
				return recursivereslice.Constructor(list)
			},
		},
		{
			name: "stack",
			skip: true,
			newIter: func(list []*nestedinteger.NestedInteger) iterator {
				return stack.Constructor(list)
			},
		},
		{
			name: "flatten-constructor",
			skip: false,
			newIter: func(list []*nestedinteger.NestedInteger) iterator {
				return flattenconstructor.Constructor(list)
			},
		},
		{
			name: "lazy-preload",
			skip: true,
			newIter: func(list []*nestedinteger.NestedInteger) iterator {
				return lazypreload.Constructor(list)
			},
		},
	}

	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				if impl.skip {
					t.Skip("not implemented")
				}
				got := collect(impl.newIter(tt.list))
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestNestedIntegerMock(t *testing.T) {
	integer := nestedinteger.NewInteger(42)
	if !integer.IsInteger() {
		t.Fatal("expected integer")
	}
	if integer.GetInteger() != 42 {
		t.Fatalf("got %d, want 42", integer.GetInteger())
	}

	empty := nestedinteger.NewEmptyList()
	if empty.IsInteger() {
		t.Fatal("empty list is not integer")
	}
	if len(empty.GetList()) != 0 {
		t.Fatalf("got len %d, want 0", len(empty.GetList()))
	}
}
