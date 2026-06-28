package main

import (
	"errors"
	"fmt"
)

type RefusalErr struct {
	reasons []error
}

func (e RefusalErr) Unwrap() []error {
	return e.reasons
}

func (e RefusalErr) Error() string {
	return fmt.Sprintf("refusing: %v", e.reasons)
}

// -----------------------------------------------------------------------

func main() {
	errWindy := errors.New("it's windy")
	errRaining := errors.New("it's raining")

	{
		// В Go 1.20+ появился способ комбинации ошибок
		//     через errors.Join().

		err := errors.Join(errRaining, errWindy)

		if errors.Is(err, errRaining) {
			fmt.Println("ouch! it's raining")
		}
		// ouch! it's raining

		if errors.Is(err, errWindy) {
			fmt.Println("ouch! it's windy")
		}
	}
	{
		err := fmt.Errorf("reasons to skip work: %w, %w",
			errRaining, errWindy)
		fmt.Println(err)
		// reasons to skip work: it's raining, it's windy
	}
	{
		err := RefusalErr{[]error{errRaining, errWindy}}
		if errors.Is(err, errRaining) {
			fmt.Println("ouch! it's raining")
		}
	}
}
