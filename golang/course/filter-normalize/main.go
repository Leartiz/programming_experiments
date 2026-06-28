package main

import (
	"fmt"
	"golang/course/filter-normalize/getters"
	"golang/course/filter-normalize/method"
)

func main() {
	(new(method.Filter)).NormalizeFilter()
	(new(method.OtherFilter)).NormalizeFilter()
	(new(method.OtherFilter1)).NormalizeFilter()

	getters.NormalizeFilter(new(getters.Filter))
	getters.NormalizeFilter(new(getters.OtherFilter))
	getters.NormalizeFilter(new(getters.OtherFilter1))

	fmt.Println("OK")
}
