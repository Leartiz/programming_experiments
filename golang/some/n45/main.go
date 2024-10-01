package main

import (
	"fmt"
)

// -----------------------------------------------------------------------

type Map[K comparable, V any] struct {
	basicMap map[K]V
}

func (m *Map[K, V]) Set(name K, value V) {
	if m.basicMap == nil {
		m.basicMap = make(map[K]V)
	}

	m.basicMap[name] = value
}

func (m *Map[K, V]) Get(name K) V {
	if m.basicMap == nil {
		m.basicMap = make(map[K]V)
	}

	return m.basicMap[name]
}

func (m *Map[K, V]) Keys() []K {
	if m.basicMap == nil {
		m.basicMap = make(map[K]V)
	}

	keys := make([]K, 0, len(m.basicMap))
	for k := range m.basicMap {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	if m.basicMap == nil {
		m.basicMap = make(map[K]V)
	}

	vals := make([]V, 0)
	for _, v := range m.basicMap {
		vals = append(vals, v)
	}
	return vals
}

// -----------------------------------------------------------------------

func main() {
	
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
