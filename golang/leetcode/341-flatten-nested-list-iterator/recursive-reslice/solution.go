package recursivereslice

// https://leetcode.com/problems/flatten-nested-list-iterator/

import "golang/leetcode/341-flatten-nested-list-iterator/nestedinteger"

type NestedIterator struct {
	root   []*nestedinteger.NestedInteger
	nested *NestedIterator
}

func Constructor(nestedList []*nestedinteger.NestedInteger) *NestedIterator {
	instance := NestedIterator{
		root:   nestedList,
		nested: nil, // !
	}
	return &instance
}

// -----------------------------------------------------------------------

func (this *NestedIterator) Next() int {
	if this.nested != nil { // если есть итератор, значит это под-список

		if this.nested.HasNext() { // соблюдаем условие использования
			return this.nested.Next()
		}

		this.nested = nil // значит, переключаемся уже на корневой список
	}

	if len(this.root) > 0 {
		top := this.root[0]
		this.root = this.root[1:] // срезаем сразу

		if top.IsInteger() {
			return top.GetInteger()
		}

		list := top.GetList()
		if len(list) > 0 {
			this.nested = Constructor(top.GetList())
			return this.Next()
		}
	}

	return 0 // значение по умолчанию!
}

func (this *NestedIterator) HasNext() bool { // prepare...?
	if this.nested == nil && len(this.root) == 0 {
		return false
	}

	// ***

	if this.nested != nil { // у последнего итератора, тут будет nil!
		if this.nested.HasNext() {
			return true
		}

		// в next будет тихий сброс
	}

	for len(this.root) > 0 {
		top := this.root[0]
		if top.IsInteger() {
			return true
		}

		if len(top.GetList()) > 0 {
			nested := Constructor(top.GetList())
			if nested.HasNext() {
				return true
			}
		}

		this.root = this.root[1:]
	}

	return false
}
