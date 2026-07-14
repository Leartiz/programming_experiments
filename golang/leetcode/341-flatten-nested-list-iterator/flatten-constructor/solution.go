package flattenconstructor

import "golang/leetcode/341-flatten-nested-list-iterator/nestedinteger"

// https://leetcode.com/problems/flatten-nested-list-iterator/

type listPos struct {
	l   []*nestedinteger.NestedInteger
	pos int
}

type NestedIterator struct {
	vals []int
	pos  int
}

func Constructor(nestedList []*nestedinteger.NestedInteger) *NestedIterator {
	vals := []int{}
	stack := []listPos{}

	rootListPos := listPos{
		l:   nestedList,
		pos: 0,
	}

	stack = append(stack, rootListPos)
	for len(stack) > 0 {
		top := stack[len(stack)-1] // копия
		idx := top.pos
		l := top.l
		for idx < len(l) {
			if l[idx].IsInteger() {
				vals = append(vals, l[idx].GetInteger())
				idx += 1
				continue
			}

			// else

			listPos := listPos{
				l:   l[idx].GetList(),
				pos: 0,
			}
			stack[len(stack)-1].pos = idx + 1
			stack = append(stack, listPos)
			break
		}

		// если обошли, удаляем
		if idx >= len(l) {
			stack = stack[:len(stack)-1] // убираем последний
		}
	}

	return &NestedIterator{
		vals: vals,
		pos:  0,
	}
}

func (this *NestedIterator) Next() int {
	actual := this.vals[this.pos]
	this.pos += 1
	return actual
}

func (this *NestedIterator) HasNext() bool {
	return this.pos < len(this.vals)
}
