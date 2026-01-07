package sol1

import "some/n5/common"

type IFilter interface {
	NormalizeFilter()
}

// -----------------------------------------------------------------------

type Filter struct {
	Limit  int
	Offset int
}

type OtherFilter struct {
	Filter
}

type OtherFilter1 struct {
	Filter

	Extra string
}

// -----------------------------------------------------------------------

func (filter *Filter) NormalizeFilter() {
	if filter.Limit < 0 {
		filter.Limit = common.Abs(filter.Limit)
		if filter.Limit > common.Instance.MAX_FILTER_LIMIT {
			filter.Limit = common.Instance.MAX_FILTER_LIMIT
		}
	}

	if filter.Offset < 0 {
		filter.Offset = common.Abs(filter.Offset)
	}
}

// ...
