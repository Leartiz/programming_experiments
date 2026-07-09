package main

import (
	"flag"
	"fmt"
)

func demoRangeOrder() {
	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
	// NOTE:
	/*
		...\map-internals> go run .
			1: 1
			2: 2
			3: 3
		...\map-internals> go run .
			1: 1
			2: 2
			3: 3
		...\map-internals> go run .
			3: 3
			1: 1
			2: 2
	*/
}

func demoMakeWrite() {
	m := make(map[int]string)
	m[1] = "1"
	m[2] = "2"
	m[3] = "3"
}

func demoNilRead() {
	var mNil map[int]string
	for k, v := range mNil {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println(mNil)
	fmt.Println(mNil[1]) // ok?
}

func demoNilWritePanic() {
	// NOTE:
	/*
		Блоки {} только ограничивают scope переменных.
		Они не ловят panic и не продолжают выполнение после него.
		recover в этой func -> после panic func завершается, дальше main идет.
	*/
	var mNil map[int]string
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic intercepted: %v\n", r)
		}
	}()

	mNil[1] = "1"
	fmt.Println("OK")
}

func demoClearVsNilVsMake() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	fmt.Println("len before clear:", len(m))

	clear(m)
	fmt.Println("len after clear:", len(m))
	m[1] = 42
	fmt.Println("reuse after clear:", m[1])

	m = make(map[int]int)
	fmt.Println("len after new make:", len(m))

	var mNil map[int]int
	fmt.Println("nil map len:", len(mNil))
}

// -----------------------------------------------------------------------

// map is not thread-safe.
//
// fatal error: concurrent map read and map write
//
// recover does NOT help - run separately: go run . -concurrent
func demoConcurrentReadWrite() {
	fmt.Println("concurrent read+write (expect fatal)...")
	m := make(map[int]int)

	go func() {
		for {
			m[1] = 1
		}
	}()

	for {
		_ = m[1]
	}
}

// -----------------------------------------------------------------------

func mutateMap(m map[string]int) {
	m["new"] = 100
	m["new_1"] = 101
	delete(m, "a")
}

func demoMapPassToFunc() {
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("before:", m)
	fmt.Println("len before:", len(m))
	mutateMap(m)
	fmt.Println("len after:", len(m))
}

// -----------------------------------------------------------------------

func main() {
	concurrent := flag.Bool("concurrent", false, "demo concurrent map fatal error")
	flag.Parse()

	if *concurrent {
		demoConcurrentReadWrite()
		return
	}

	fmt.Println("--- range order ---")
	demoRangeOrder()

	fmt.Println("--- make write ---")
	demoMakeWrite()

	fmt.Println("--- nil read ---")
	demoNilRead()

	fmt.Println("--- nil write panic ---")
	demoNilWritePanic()

	fmt.Println("--- clear vs nil vs make ---")
	demoClearVsNilVsMake()

	fmt.Println("--- map pass to func ---")
	demoMapPassToFunc()

	fmt.Println("tip: go run . -concurrent")
}
