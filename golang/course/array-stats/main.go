package main

import "fmt"

func generateArray1(length int) []int {
	result := make([]int, 0, length)
	for i := length; i > 0; i-- {
		result = append(result, i*3)
	}
	return result
}

func generateArray2(length int) []int {
	result := make([]int, 0, length)
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			result = append(result, 1)
		} else {
			result = append(result, i%3)
		}
	}
	return result
}

func generateArray3(length int) []int {
	result := make([]int, 0, length*2)
	for i := 1; i <= length; i++ {
		result = append(result, i)
	}
	for i := length; i > 0; i-- {
		result = append(result, i)
	}
	return result
}

func countNumbers(data []int) string {
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

	return fmt.Sprintf("%v, %v, %v, %v, %v",
		uniqueCount, duplicateCount, zeroCount, mod2Eq0Count, mod2Eq1Count)
}

func main() {
	arr := []int{1, 2, 3, 1, 2, 5, 4, 1, 2, 0, 0, 0}
	fmt.Println(countNumbers(arr))
}
