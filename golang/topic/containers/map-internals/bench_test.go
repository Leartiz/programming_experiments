package main

import (
	"sync"
	"testing"
)

// go test -bench=. -benchmem
// -----------------------------------------------------------------------

var benchSink any

// read-heavy: 90% Load, 10% Store
// go test -bench=ReadHeavy -benchmem
// -----------------------------------------------------------------------

func BenchmarkMapMutex_ReadHeavy(b *testing.B) {
	var mu sync.RWMutex
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		m[i] = i
	}

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 100
			if i%10 == 0 { // 10% Store
				mu.Lock()
				m[key] = i
				mu.Unlock()
			} else { // 90% Load
				mu.RLock()
				benchSink = m[key]
				mu.RUnlock()
			}
			i++
		}
	})
}

func BenchmarkSyncMap_ReadHeavy(b *testing.B) {
	var sm sync.Map
	for i := 0; i < 100; i++ {
		sm.Store(i, i)
	}

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 100
			if i%10 == 0 {
				sm.Store(key, i)
			} else {
				benchSink, _ = sm.Load(key)
			}
			i++
		}
	})
}

// write-heavy: 50% Load, 50% Store
// go test -bench=WriteHeavy -benchmem
// -----------------------------------------------------------------------

func BenchmarkMapMutex_WriteHeavy(b *testing.B) {
	var mu sync.RWMutex
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		m[i] = i
	}

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 100
			if i%2 == 0 {
				mu.Lock()
				m[key] = i
				mu.Unlock()
			} else {
				mu.RLock()
				benchSink = m[key]
				mu.RUnlock()
			}
			i++
		}
	})
}

func BenchmarkSyncMap_WriteHeavy(b *testing.B) {
	var sm sync.Map
	for i := 0; i < 100; i++ {
		sm.Store(i, i)
	}

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 100
			if i%2 == 0 {
				sm.Store(key, i)
			} else {
				benchSink, _ = sm.Load(key)
			}
			i++
		}
	})
}

// NOTE:
/*
	BenchmarkMapMutex_ReadHeavy-16          28081341                41.92 ns/op            0 B/op          0 allocs/op
	BenchmarkSyncMap_ReadHeavy-16           134834110               10.31 ns/op            5 B/op          0 allocs/op

	BenchmarkMapMutex_WriteHeavy-16         36993079                30.66 ns/op            0 B/op          0 allocs/op
	BenchmarkSyncMap_WriteHeavy-16          74904028                17.49 ns/op           28 B/op          0 allocs/op
*/
