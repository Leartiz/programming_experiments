package lazypreload

// https://leetcode.com/problems/flatten-nested-list-iterator/

import "golang/leetcode/341-flatten-nested-list-iterator/nestedinteger"

type NestedIterator struct {
	// TODO: queue/stack + prepare next integer on HasNext
}

func Constructor(nestedList []*nestedinteger.NestedInteger) *NestedIterator {
	return &NestedIterator{}
}

func (it *NestedIterator) Next() int {
	return 0
}

func (it *NestedIterator) HasNext() bool {
	return false
}
