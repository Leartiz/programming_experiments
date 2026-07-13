package nestedinteger

// Mock LeetCode API for local run and tests.

type NestedInteger struct {
	val  *int
	list []*NestedInteger
}

func NewInteger(v int) *NestedInteger {
	return &NestedInteger{val: &v}
}

func NewList(list []*NestedInteger) *NestedInteger {
	if list == nil {
		list = []*NestedInteger{}
	}
	return &NestedInteger{list: list}
}

func NewEmptyList() *NestedInteger {
	return NewList([]*NestedInteger{})
}

func (n *NestedInteger) IsInteger() bool {
	return n.val != nil
}

func (n *NestedInteger) GetInteger() int {
	return *n.val
}

func (n *NestedInteger) GetList() []*NestedInteger {
	return n.list
}
