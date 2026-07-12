package main

import (
	"fmt"

	"golang/leetcode/41-first-missing-positive/canonical"
	fixedarray "golang/leetcode/41-first-missing-positive/fixed-array"
)

func main() {
	fmt.Println("fixed-array:")
	{
		nums := []int{1, 2, 0}
		fmt.Println(fixedarray.FirstMissingPositive(nums))
	}
	{
		nums := []int{3, 4, -1, 1}
		fmt.Println(fixedarray.FirstMissingPositive(nums))
	}
	{
		nums := []int{7, 8, 9, 11, 12}
		fmt.Println(fixedarray.FirstMissingPositive(nums))
	}

	fmt.Println("canonical:")
	{
		nums := []int{1, 2, 0}
		fmt.Println(canonical.FirstMissingPositive(nums))
	}
	{
		nums := []int{3, 4, -1, 1}
		fmt.Println(canonical.FirstMissingPositive(nums))
	}
	{
		nums := []int{7, 8, 9, 11, 12}
		fmt.Println(canonical.FirstMissingPositive(nums))
	}
}
