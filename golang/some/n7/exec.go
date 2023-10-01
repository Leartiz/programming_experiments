package n7

import (
	"encoding/json"
	"fmt"

	"github.com/markphelps/optional"
)

func Exec() {
	var value = &struct {
		Field optional.String `json:"field"`
	}{}

	err := json.Unmarshal([]byte(`{"field":"bar"}`), value)
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
