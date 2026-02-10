package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]CacheEntry
	mu   sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]CacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.data {
			if time.Since(value.createdAt) > interval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
