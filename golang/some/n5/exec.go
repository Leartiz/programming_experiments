package n5

import (
	"fmt"
	"some/n5/sol1"
	"some/n5/sol2"
)

func Exec() {
	fmt.Println("n5")

	// sol1
	{
		{
			var ff = new(sol1.Filter)
			ff.NormalizeFilter()
		}
		{
			var ff = new(sol1.OtherFilter)
			ff.NormalizeFilter()
		}
		{
			var ff = new(sol1.OtherFilter1)
			ff.NormalizeFilter()
		}
	}

	// ***

	// sol2
	{
		{
			var ff = new(sol2.Filter)
			sol2.NormalizeFilter(ff)
		}
		{
			var ff = new(sol2.OtherFilter)
			sol2.NormalizeFilter(ff)
		}
		{
			var ff = new(sol2.OtherFilter1)
			sol2.NormalizeFilter(ff)
		}
	}
}
