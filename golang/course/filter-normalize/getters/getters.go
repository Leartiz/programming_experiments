package getters

import "golang/course/filter-normalize/common"

type IFilter interface {
	Limit() int
	Offset() int
	SetLimit(int)
	SetOffset(int)
}

type Filter struct {
	limit  int
	offset int
}

type OtherFilter struct {
	Filter
}

type OtherFilter1 struct {
	Filter
	Extra string
}

func (ff *Filter) Limit() int        { return ff.limit }
func (ff *Filter) Offset() int       { return ff.offset }
func (ff *Filter) SetLimit(val int)  { ff.limit = val }
func (ff *Filter) SetOffset(val int) { ff.offset = val }

func NormalizeFilter(filter IFilter) {
	if filter.Limit() < 0 {
		filter.SetLimit(common.Abs(filter.Limit()))
		if filter.Limit() > common.Instance.MaxFilterLimit {
			filter.SetLimit(common.Instance.MaxFilterLimit)
		}
	}
	if filter.Offset() < 0 {
		filter.SetOffset(common.Abs(filter.Offset()))
	}
}
