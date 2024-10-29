package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {

	fmt.Println("\n*** Slices ***")

	{
		names := []string{"bob", "alice", "cindy"}
		fmt.Printf("names: %v\n", names)

		slices.Sort(names)
		fmt.Printf("names: %v\n", names)

		n, found := slices.BinarySearch(names, "bob")
		fmt.Println(n, found)
	}

	{
		nums1 := []int{1, 2}
		nums2 := []int{2, 3}
		nums3 := []int{4, 5}
		nums := slices.Concat(nums1, nums2, nums3)
		fmt.Printf("concated nums: %v\n", nums)
	}

	{
		names := []string{"alice", "bob", "cindy", "dave"}
		names = slices.Delete(names, 1, 3) // [..)
		fmt.Printf("names after deleted vals: %v\n", names)
	}

	{
		nums1 := []int{1, 2, 3}
		nums2 := []int{1, 2, 3}

		fmt.Println(slices.Equal(nums1, nums2)) // true
		nums3 := []int{3, 2, 1}

		// Порядок элементов имеет значение!
		fmt.Println(slices.Equal(nums1, nums3)) // false
	}

	{
		names := []string{"alice", "bob", "cindy"}
		idx := slices.Index(names, "bob")
		fmt.Printf("value index: %v\n", idx) // 1
	}

	{
		names := []string{"alice", "bob", "cindy"}
		fmt.Printf("names: %v\n", names)

		slices.Reverse(names)
		fmt.Printf("reversed names: %v\n", names)
	}

	fmt.Println("\n*** Maps ***")

	{
		m1 := map[string]int{"one": 1, "two": 2}
		fmt.Printf("m1: %v\n", m1)

		m2 := maps.Clone(m1)
		fmt.Printf("m2: %v\n", m2)
	}

	{
		m := map[string]int{
			"one": 1,
			"two": 2,
			"thr": 3,
		}
		maps.DeleteFunc(m, func(key string, val int) bool {
			return val%2 == 0
		})
		fmt.Println(m)
	}

	{
		m1 := map[string]int{"one": 1, "two": 2}
		m2 := map[string]int{"two": 2, "one": 1}
		fmt.Println(maps.Equal(m1, m2))
	}
}
