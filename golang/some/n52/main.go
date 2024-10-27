package main

import (
	"fmt"
)

func main() {
	{
		arr := [5]int{11, 22, 33}
		// [11 22 33 0 0]

		s := arr[1:3]
		// [22 33]

		// Емкость среза равна cap(src)-low.

		fmt.Printf("len: %v\n", len(s))
		fmt.Printf("cap: %v\n", cap(s))

		s = append(s, 10)
		fmt.Printf("arr: %v\n", arr)
		fmt.Printf("s: %v\n", s)
	}

	{
		arr := [5]int{11, 22, 33}
		// [11 22 33 0 0]

		s := arr[1:3:3] // !!!
		// [22 33]

		// Полносрезное выражение работает почти как обычный срез,
		// только емкость у результата будет не cap(src)-low, а max-low.

		fmt.Printf("len: %v\n", len(s))
		fmt.Printf("cap: %v\n", cap(s))

		s = append(s, 10)
		fmt.Printf("arr: %v\n", arr)

		// s больше не использует arr
		fmt.Printf("s: %v\n", s)
	}
}
