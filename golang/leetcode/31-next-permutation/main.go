package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/next-permutation/

func printInts(nums []int) {
	for _, num := range nums {
		fmt.Printf("%v ", num)
	}
	fmt.Println()
}

func nextPermutation(nums []int) {
	// NOTE:
	/*
		Constraints:
			1 <= nums.length <= 100
			0 <= nums[i] <= 100

		The replacement must be in place and
		use only constant extra memory.
	*/

	// NOTE:
	/*
		arr = [1,2,3] is [1,3,2]
		arr = [2,3,1] is [3,1,2]
		arr = [3,2,1] is [1,2,3]

		Алгоритм Нарайаны
	*/

	pivotIndex := -1
	for i := len(nums) - 2; i >= 0; i -= 1 {
		if nums[i] < nums[i+1] {
			pivotIndex = i
			break
		}
	}

	// если "спада" нет, массив отсортирован по убывания
	if pivotIndex == -1 {
		slices.Reverse(nums)
		return
	}

	// ищем первый элемент, который больше "опорного" (на замену)
	for i := len(nums) - 1; i > pivotIndex; i -= 1 {
		if nums[i] > nums[pivotIndex] {

			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
			slices.Reverse(nums[pivotIndex+1:])
			break
		}
	}
}

func main() {
	{
		nums := []int{2, 3, 1}
		printInts(nums)

		nextPermutation(nums)
		printInts(nums)
	}
	{
		nums := []int{2, 2, 0, 4, 3, 1}
		printInts(nums)

		nextPermutation(nums)
		printInts(nums)
	}
}
