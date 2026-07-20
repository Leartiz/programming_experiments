package main

import (
	"sync"
	"testing"
)

// go test "-bench=." "-cpu=1,4"

func BenchmarkSpinLock_Uncontended(b *testing.B) {
	var mu spinLock
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

func BenchmarkMutex_Uncontended(b *testing.B) {
	var mu sync.Mutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

// -----------------------------------------------------------------------

func BenchmarkSpinLock_Contended(b *testing.B) {
	var (
		mu spinLock
		n  int
	)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			n++
			mu.Unlock()
		}
	})
}

func BenchmarkMutex_Contended(b *testing.B) {
	var (
		mu sync.Mutex
		n  int
	)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			n++
			mu.Unlock()
		}
	})
}
