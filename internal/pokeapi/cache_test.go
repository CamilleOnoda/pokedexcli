package pokeapi

import (
	"context"
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	const ttl = 2 * time.Second
	tests := []struct {
		name  string
		key   string
		value []byte
	}{
		{
			name:  "simple set and get",
			key:   "https://example.com",
			value: []byte("test data"),
		},
		{
			name:  "another set and get",
			key:   "https://example.com/path",
			value: []byte("more test data"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			cleanupInterval := 100 * time.Millisecond
			cache := NewCache(ctx, cleanupInterval)
			cache.Set(test.key, test.value, ttl)
			cachedValue, found := cache.Get(test.key)
			if !found {
				t.Errorf("Expected to find key '%s' in cache", test.key)
				return
			}
			if string(cachedValue) != string(test.value) {
				t.Errorf("Expected value '%s', got '%s'", string(test.value), string(cachedValue))
				return
			}
		})
	}
}

func TestGetExpired(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		value        []byte
		ttl          time.Duration
		shouldExpire bool
	}{
		{
			name:         "remove expired entry",
			key:          "https://example.com/expire",
			value:        []byte("expired data"),
			ttl:          5 * time.Millisecond,
			shouldExpire: true,
		},

		{
			name:         "unexpired entry",
			key:          "https://example.com/unexpire",
			value:        []byte("unexpired data"),
			ttl:          1 * time.Second,
			shouldExpire: false,
		},
	}

	const waitTime = 10 * time.Millisecond

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			cleanupInterval := 100 * time.Millisecond
			cache := NewCache(ctx, cleanupInterval)
			cache.Set(test.key, test.value, test.ttl)

			cachedValue, found := cache.Get(test.key)
			if !found {
				t.Errorf("Expected to find key'%s' in cache", test.key)
				return
			} else if string(cachedValue) != string(test.value) {
				t.Errorf("Expected value '%s', got '%s'", string(test.value), string(cachedValue))
				return
			}

			time.Sleep(waitTime)

			cachedValue, found = cache.Get(test.key)

			if test.shouldExpire && found {
				t.Errorf("Expected key '%s' to be expired", test.key)
			}
			if !test.shouldExpire && !found {
				t.Errorf("Expected key '%s' to be still present", test.key)
			}
		})
	}
}
