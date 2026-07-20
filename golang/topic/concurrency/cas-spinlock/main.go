package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// spinlock: только CAS, без сна
// при contention (спор) горутина крутит CPU в цикле

type spinLock struct {
	locked int32 // 0 free, 1 held
}

func (s *spinLock) Lock() {
	// NOTE:
	//   CAS - не обычный Go if; под капотом CPU-инструкция.
	//   amd64:  mov eax, 0 ; lock cmpxchg [locked], 1
	//   arm64:  ldxr / cmp / stxr  (цикл load-exclusive / store-exclusive)
	//
	for !atomic.CompareAndSwapInt32(&s.locked, 0, 1) {
		// spin
	}
}

func (s *spinLock) Unlock() {
	atomic.StoreInt32(&s.locked, 0)
}

func main() {
	var (
		mu spinLock
		n  int
		wg sync.WaitGroup
	)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				mu.Lock()
				fmt.Println("hold", id)
				n++
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("n =", n) // 20
}
