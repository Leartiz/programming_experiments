package cache

import (
	"sync"
	"sync/atomic"
	"time"
)

// Идеи:
//   Value interface{} ---> generics (хотим ли?)
//

type Item struct {
	Value      interface{}
	Expiration time.Time
}

// Cache - простой in-memory кэш с TTL.
type Cache struct {
	mu    sync.RWMutex
	items map[string]Item

	cleanupTicker *time.Ticker
	cleanupDone   chan struct{}
	cleanupWg     sync.WaitGroup

	stopped atomic.Bool
}

// NewCache создает новый кэш и запускает фоновую очистку.
func NewCache(cleanupInterval time.Duration) *Cache {

	// cleanupInterval, можно конечно сделать error,
	// НО и допустимо, чтобы NewTicker кинул панику.

	c := &Cache{
		mu:    sync.RWMutex{},
		items: make(map[string]Item),

		cleanupTicker: time.NewTicker(cleanupInterval),
		cleanupDone:   make(chan struct{}),
		cleanupWg:     sync.WaitGroup{},

		stopped: atomic.Bool{},
	}

	c.cleanupWg.Add(1)
	go func() {
		defer c.cleanupWg.Done()
		c.cleanupLoop()
	}()

	return c
}

// тут, к вопросу как поступить.
// разрешить ли Get/Set/... после остановки?
func (c *Cache) Shutdown() {
	if !c.stopped.CompareAndSwap(false, true) {
		// уже остановлен!
		return
	}

	// After Stop, no more ticks will be sent.
	// Stop does not close the channel,
	// to prevent a concurrent goroutine reading from the channel
	// from seeing an erroneous "tick".
	//
	c.cleanupTicker.Stop()
	close(c.cleanupDone) // завершить горутину очистки
	c.cleanupWg.Wait()   // ждем завершения горутины!

	c.mu.Lock()
	defer c.mu.Unlock()

	for k := range c.items {
		delete(c.items, k)
	}
}

// -----------------------------------------------------------------------

// Set добавляет или обновляет значение в кэше.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {

	// ttl = 0, ОК пусть допустимо...
	// key = "", ОК пусть будет
	// value = nil, ХМ, может так надо?

	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
}

// Get возвращает значение, если оно есть и не истекло.
func (c *Cache) Get(key string) (interface{}, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if time.Now().After(item.Expiration) {
		// удалять не будем!
		// пусть фоновая задача крутит и чистит.

		return nil, false
	}

	return item.Value, true
}

// Delete удаляет ключ из кэша.
func (c *Cache) Delete(key string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// cleanupLoop периодически удаляет истекшие элементы.
func (c *Cache) cleanupLoop() {
	for {
		select {
		case <-c.cleanupDone:
			return

		case <-c.cleanupTicker.C:
			c.mu.Lock()
			for key, item := range c.items {
				if time.Now().After(item.Expiration) {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

// Size возвращает количество элементов в кэше (включая истекшие, но еще не удаленные).
func (c *Cache) Size() int {

	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// NOTE:
/*
	type Item struct {
		Value      interface{}
		Expiration time.Time
	}

	// Cache - простой in-memory кэш с TTL.
	type Cache struct {
		mu    sync.Mutex
		items map[string]Item
	}

	// NewCache создает новый кэш и запускает фоновую очистку.
	func NewCache(cleanupInterval time.Duration) *Cache {
		c := &Cache{
			items: make(map[string]Item),
		}

		go c.cleanupLoop(cleanupInterval)

		return c
	}

	// Set добавляет или обновляет значение в кэше.
	func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.items[key] = Item{
			Value:      value,
			Expiration: time.Now().Add(ttl),
		}
	}

	// Get возвращает значение, если оно есть и не истекло.
	func (c *Cache) Get(key string) (interface{}, bool) {
		c.mu.Lock()
		defer c.mu.Unlock()

		item, found := c.items[key]
		if !found {
			return nil, false
		}
		if time.Now().After(item.Expiration) {
			delete(c.items, key)
			return nil, false
		}
		return item.Value, true
	}

	// Delete удаляет ключ из кэша.
	func (c *Cache) Delete(key string) {
		c.mu.Lock()
		defer c.mu.Unlock()

		delete(c.items, key)
	}

	// cleanupLoop периодически удаляет истекшие элементы.
	func (c *Cache) cleanupLoop(interval time.Duration) {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			c.mu.Lock()
			for key, item := range c.items {
				if time.Now().After(item.Expiration) {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}

	// Size возвращает количество элементов в кэше (включая истекшие, но еще не удаленные).
	func (c *Cache) Size() int {
		c.mu.Lock()
		defer c.mu.Unlock()

		return len(c.items)
	}
*/
