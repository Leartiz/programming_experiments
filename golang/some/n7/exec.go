package n7

import (
	"encoding/json"
	"fmt"

	"github.com/markphelps/optional"
)

func Exec() {
	fmt.Println("--- 1 ---")

	{
		var value = &struct {
			Field optional.String `json:"field"`
		}{}

		err := json.Unmarshal([]byte(
			`{ "field": "bar" }`,
		), value)

		if err != nil {
			fmt.Println(err)
		}

		value.Field.If(func(s string) {
			fmt.Println(s)
		})

		_, err = value.Field.Get()
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("--- 2 ---")

	{
		var value = &struct {
			Field optional.String `json:"field"`
		}{}

		err := json.Unmarshal([]byte(
			`{}`,
		), value)

		if err != nil {
			fmt.Println(err)
		}

		value.Field.If(func(s string) {
			fmt.Println(s)
		})

		_, err = value.Field.Get()
		if err != nil {
			fmt.Println(err) // value not present
		}
	}
}
