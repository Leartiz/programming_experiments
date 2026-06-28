package main

import (
	"fmt"
)

type AssertionWithParams struct {
	AssertFunc func(params ...any) error
	Params     []any
}

func assert1(params ...any) error {
	/* val int */

	val := params[0].(int)

	fmt.Printf("assert1(%v)\n", val)
	return nil
}

func assert2(params ...any) error {
	/* p, p1 string */

	p := params[0].(string)
	p1 := params[1].(string)

	fmt.Printf("assert2(%v, %v)\n", p, p1)
	return nil
}

func checkAssertions(list []AssertionWithParams) error {
	for _, val := range list {
		if err := val.AssertFunc(val.Params...); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	assertions := []AssertionWithParams{{
		AssertFunc: assert1,
		Params:     []any{1},
	}, {
		AssertFunc: assert2,
		Params:     []any{"str0", "str1"},
	}}
	//...

	err := checkAssertions(assertions)
	fmt.Println(err)
}
