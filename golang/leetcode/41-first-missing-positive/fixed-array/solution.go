package fixedarray

// https://leetcode.com/problems/first-missing-positive/

func FirstMissingPositive(nums []int) int {
	// 1 <= nums.length <= 105
	// 10^5 = 100'000
	// завести какой-то массив вот такой фиксированной длины? а зачем?
	// найти минимальный, позитивный, он будет нулевым индексом в массиве этом!

	// сортировать мы не можем!
	// память дополнительная O(1)

	// попробовать найти минимальный?
	// и

	// [3,4,-1,1]
	// 1

	// посчитать количество позитивных чисел? и разделить?
	// 3, 1
	// сумму позитивных 1 3 4 = 8
	//

	minPositive := 0 // старовая позиция, по сути
	for _, num := range nums {
		if num > 0 {
			if minPositive == 0 {
				minPositive = num
			} else if minPositive > num {
				minPositive = num
			}
		}
	}

	// быстрый выход, по сути число первое
	if minPositive == 0 || minPositive > 1 {
		return 1
	}

	const maxNumsSize = 100001
	fixed := [maxNumsSize]bool{}
	fixed[0] = true // index = minPositive - minPositive
	for _, num := range nums {
		if num > minPositive {
			index := num - minPositive
			if index < maxNumsSize {
				fixed[index] = true
			}
		}
	}

	for i, f := range fixed {
		if !f {
			return i + minPositive
		}
	}

	return -1 // по условию не возможно (хитрим с размером O(1)) ;)
}
