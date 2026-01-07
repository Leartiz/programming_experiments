package n8

import (
	"encoding/json"
	"fmt"
)

func Exec() {
	fmt.Println("--- 1 ---")

	{
		var value = struct {
			Field *string `json:"field,omitempty"`
		}{}

		out, _ := json.Marshal(value)
		fmt.Println(string(out))
	}

	fmt.Println("--- 2 ---")

	{
		var value = struct {
			Field *string `json:"field"`
		}{}

		out, _ := json.Marshal(value)
		fmt.Println(string(out))
	}
}
