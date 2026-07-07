package main

import "fmt"

// https://leetcode.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
	// 1 <= target <= 10^9

	windowSize := 1
	for windowSize <= len(nums) {
		curSum := 0

		// сумма первого окна
		for i := 0; i < windowSize; i += 1 {
			curSum += nums[i]
		}
		if curSum >= target {
			return windowSize
		}

		// следующие суммы считаем по принципу +следующий, -предидущий
		/*
		   0 1 2 3 4 5
		   2,3,1,2,4,3
		   windowSize = 2
		   firstWindowSize = 2 + 3

		   и потом ...

		   ---

		   0 1 2 3 4
		   1,2,3,4,5

		   5 - 3 = 2

		*/
		for i := 0; i < len(nums)-windowSize; i += 1 {
			curSum -= nums[i]
			curSum += nums[i+windowSize] // добавить следующий

			if curSum >= target {
				return windowSize
			}
		}

		windowSize += 1
	}

	// If there is no such subarray, return 0 instead.
	return 0
}

func main() {
	{
		nums := []int{1, 2, 3, 4, 5}
		fmt.Printf("%v\n", minSubArrayLen(15, nums))
	}
}
