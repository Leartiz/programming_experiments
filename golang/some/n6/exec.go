package n6

import "fmt"

type IFilter interface {
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

// -----------------------------------------------------------------------

func Exec() {
	var f IFilter = OtherFilter{}

	switch f.(type) {
	case Filter:
		fmt.Println("is Filter")
	case OtherFilter:
		fmt.Println("is OtherFilter")
	case OtherFilter1:
		fmt.Println("is OtherFilter1")
	}

	fmt.Println("[END]")
}
