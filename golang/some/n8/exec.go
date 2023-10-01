package n8

import (
	"encoding/json"
	"fmt"
)

func Exec() {
	var value = struct {
		Field *string `json:"field,omitempty"`
	}{}

	out, _ := json.Marshal(value)

	fmt.Println(string(out))
	// outputs: {"field":"bar"}
}
