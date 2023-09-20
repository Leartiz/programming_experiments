package sol3

import "some/n5/common"

type Filter struct {
	Limit  int
	Offset int
}

func (f *Filter) GetLimit() int {
	return f.Limit
}

type IFilter interface {
	~struct {
		Limit  int
		Offset int
	}
}

func NormalizeFilter(filter IFilter) {
	if filter.Limit < 0 {
		filter.SetLimit(common.Abs(filter.Limit()))
		if filter.Limit() > common.Instance.MAX_FILTER_LIMIT {
			filter.SetLimit(common.Instance.MAX_FILTER_LIMIT)
		}
	}

	if filter.Offset() < 0 {
		filter.SetOffset(common.Abs(filter.Offset()))
	}
}

// type Foo struct {
// 	val int
// }

// func (f Foo) GetVal() int {
// 	return f.val
// }

// type Bar struct {
// 	val int
// }

// func (b Bar) GetVal() int {
// 	return b.val
// }

// type MyType interface {
// 	Foo | Bar
// 	GetVal() int
// }

// func Add[T MyType](slice []T) int {
// 	var sum int
// 	for _, elem := range slice {
// 		sum += elem.GetVal()
// 	}
// 	return sum
// }
