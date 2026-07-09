package main

import (
	"fmt"

	"golang/leetcode/2166-design-bitset/chunkmap"
)

func main() {
	obj := chunkmap.Constructor(5)
	obj.Fix(3)
	obj.Fix(1)
	obj.Flip()
	fmt.Println(obj.All(), obj.One(), obj.Count(), obj.ToString())
}
