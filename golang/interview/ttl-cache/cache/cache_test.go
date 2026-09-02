package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_CAS(t *testing.T) {
	stopped := atomic.Bool{}
	if stopped.CompareAndSwap(false, true) {
		fmt.Println("false true")
	}
	if stopped.CompareAndSwap(false, true) {
		fmt.Println("false true")
	}

	if stopped.CompareAndSwap(true, false) {
		fmt.Println("true false")
	}
}

func Benchmark_Get(b *testing.B) {
	const benchKey = "bench-key"

	c := NewCache(5)
	c.Set(benchKey, "string", 10)

	for i := 0; i < b.N; i += 1 {
		c.Get(benchKey)
	}

	c.Shutdown()
}

func TestConcurrentAccess(t *testing.T) {
	c := NewCache(time.Millisecond * 50)
	defer c.Shutdown()

	// ***

	const goroutines = 50
	const iterations = 1000

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done() // -1

			for j := 0; j < iterations; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j%10)

				c.Set(key, j, time.Millisecond*10)
				_, _ = c.Get(key)

				// случайно удалить
				if j%3 == 0 {
					c.Delete(key)
				}
			}
		}(i)
	}

	wg.Wait()
}

func BenchmarkSetGet(b *testing.B) {
	c := NewCache(time.Minute)
	defer c.Shutdown()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set("key", "value", time.Minute)
			_, _ = c.Get("key")
		}
	})
}

// go test -fuzz=FuzzSetGet -fuzztime=1s
// go test -v -fuzz=FuzzSetGet -fuzztime=1s
func FuzzSetGet(f *testing.F) {
	f.Add("key", "value", 1000)
	f.Fuzz(func(t *testing.T, key string, value string, ttl int) {
		c := NewCache(time.Millisecond)
		defer c.Shutdown()
		if ttl < 0 {
			ttl = 0
		}

		c.Set(key, value, time.Duration(ttl))
		_, _ = c.Get(key)
	})
}
