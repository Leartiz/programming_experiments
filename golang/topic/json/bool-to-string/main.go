package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func booleanToString(data []interface{}) []interface{} {
	// [true,false,1,"hello",false,true]

	convertedData := make([]interface{}, 0)
	for i := range data {
		switch data[i].(type) {
		case bool:
			value := data[i].(bool)
			if value {
				convertedData = append(convertedData, "on")
			} else {
				convertedData = append(convertedData, "off")
			}
		default:
			convertedData = append(convertedData, data[i])
		}
	}
	return convertedData
}

func main() {
	data := ReadInput()
	result, _ := json.Marshal(booleanToString(data))
	fmt.Println(string(result))
}

func ReadInput() []interface{} {
	var data []interface{}
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
