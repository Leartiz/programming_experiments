package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func ZipMap[Key comparable, Val any](
	keys []Key, vals []Val) map[Key]Val {

	resultMap := map[Key]Val{}
	maxLen := min(len(keys), len(vals))
	for i := 0; i < maxLen; i++ {
		resultMap[keys[i]] = vals[i]
	}
	return resultMap
}

func main() {
	{
		a, b := 5, 3
		max := Max(a, b)
		fmt.Println(max) // 5
	}
	{
		a, b := 5.5, 3.3
		max := Max(a, b)
		fmt.Println(max) // 5.5
	}

	{
		m := ZipMap([]string{"one"}, []int{11, 22, 33})
		fmt.Println(m)
		// map[one:11]
	}
	{
		keys := []string{"one", "two", "thr"}
		vals := []int{11, 22, 33}

		m := ZipMap(keys, vals)
		fmt.Println(m)
		// map[one:11 two:22 thr:33]
	}
}
