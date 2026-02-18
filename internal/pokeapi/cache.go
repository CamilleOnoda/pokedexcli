package pokeapi

import (
	"sync"
	"time"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{}, ttl time.Duration)
	Clear()
	cleanupExpired()
	removeEpired()
}

type cacheEntry struct {
	expiresAt time.Time
	value     interface{}
}

type InMemoryCache struct {
	data map[string]cacheEntry
	mu   sync.RWMutex
}

func NewCache() *InMemoryCache {
	cache := &InMemoryCache{
		data: make(map[string]cacheEntry),
	}
	go cache.cleanupExpired()
	return cache
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return val.value, true
}

func (c *InMemoryCache) Set(key string, val interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		expiresAt: time.Now().Add(ttl),
		value:     val,
	}

}

func (c *InMemoryCache) cleanupExpired() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.removeEpired()
	}
}

func (c *InMemoryCache) removeEpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.data {
		if now.After(entry.expiresAt) {
			delete(c.data, key)
		}
	}
}

func (c *InMemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string]cacheEntry)
}
