package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func findFirstNegative(data []int) int {
	firstNegative := 0
	for i := range data {
		if data[i] < 0 {
			firstNegative = data[i]
			return firstNegative
		}
	}
	return firstNegative
}

func main() {
	data := ReadInput()
	result := findFirstNegative(data)
	fmt.Println(result)
}

func ReadInput() []int {
	var data []int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	json.Unmarshal([]byte(input), &data)
	return data
}
