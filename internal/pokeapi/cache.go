package pokeapi

// Stores data with expiration (TTL)

import (
	"context"
	"sync"
	"time"
)

type Cache interface {
	Get(key string) ([]byte, bool)
	Set(key string, val []byte, ttl time.Duration)
	Clear()
}

type cacheEntry struct {
	expiresAt time.Time
	value     []byte
}

type InMemoryCache struct {
	data map[string]cacheEntry
	mu   sync.RWMutex
}

func NewCache(ctx context.Context, cleanupInterval time.Duration) *InMemoryCache {
	cache := &InMemoryCache{
		data: make(map[string]cacheEntry),
	}
	go cache.cleanupExpired(ctx, cleanupInterval)
	return cache
}

func (c *InMemoryCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	if !ok || time.Now().After(entry.expiresAt) {
		return nil, false
	}
	return entry.value, true
}

func (c *InMemoryCache) Set(key string, val []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		expiresAt: time.Now().Add(ttl),
		value:     val,
	}

}

func (c *InMemoryCache) cleanupExpired(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.removeExpired()
		case <-ctx.Done():
			return
		}
	}
}

func (c *InMemoryCache) removeExpired() {
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
