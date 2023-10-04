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

func countNumbers(data []int) string {
	/*
	   уникальных чисел
	   дублированных чисел
	   нулей
	   четных чисел
	   нечетных чисел
	*/

	zeroCount := 0
	mod2Eq0Count := 0
	mod2Eq1Count := 0
	mp := make(map[int]int)
	for _, val := range data {
		if val == 0 {
			zeroCount++
		}

		if val != 0 {
			if val%2 == 0 {
				mod2Eq0Count++
			} else {
				mod2Eq1Count++
			}
		}

		mp[val]++
	}

	uniqueCount := 0
	duplicateCount := 0

	for _, val := range mp {
		if val == 1 {
			uniqueCount++
		} else {
			duplicateCount++
		}
	}

	return fmt.Sprintf(
		"%v, %v, %v, %v, %v",
		uniqueCount, duplicateCount,
		zeroCount,
		mod2Eq0Count, mod2Eq1Count,
	)
}

func Exec() {
	arr := []int{1, 2, 3, 1, 2, 5, 4, 1, 2, 0, 0, 0}
	fmt.Println(
		countNumbers(
			arr))
}
