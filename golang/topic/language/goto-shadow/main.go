package main

import (
	"fmt"
)

type Info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
		}
	}

	fmt.Println("out!")

	// ***

	var info Info
	var err error = nil
	info.result, err = work()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", info)

	// ***

	randomShadowVar()
	capFromMap()
}

func randomShadowVar() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)

		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}

func capFromMap() {
	m := make(map[string]int, 99)
	fmt.Printf("len: %v", len(m))

	//fmt.Printf("cap: %v", cap(m))
}
