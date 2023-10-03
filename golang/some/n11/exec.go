package n11

import "fmt"

func generateArray_1(length int) []int {
	result := make([]int, 0, length)

	var i = length
	for i > 0 {
		result = append(result, i*3)
		i--
	}
	return result
}

func generateArray_2(length int) []int {
	result := make([]int, 0, length)
	var i = 0
	for i < length {
		if i%2 == 0 {
			result = append(result, 1)
		} else {
			result = append(result, i%3)
		}
		i++
	}
	return result
}

func generateArray_3(length int) []int {
	result := make([]int, 0, length*2)
	i := 1
	for i <= length {
		result = append(result, i)
		i++
	}
	i = length
	for i > 0 {
		result = append(result, i)
		i--
	}
	return result
}

func Exec() {
	fmt.Println(
		generateArray_3(
			5))
}
