package method

import "golang/course/filter-normalize/common"

type IFilter interface {
	NormalizeFilter()
}

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

func (filter *Filter) NormalizeFilter() {
	if filter.Limit < 0 {
		filter.Limit = common.Abs(filter.Limit)
		if filter.Limit > common.Instance.MaxFilterLimit {
			filter.Limit = common.Instance.MaxFilterLimit
		}
	}
	if filter.Offset < 0 {
		filter.Offset = common.Abs(filter.Offset)
	}
}
