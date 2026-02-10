package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	tests := []struct {
		name  string
		key   string
		value []byte
	}{
		{
			name:  "simple add and get",
			key:   "https://example.com",
			value: []byte("test data"),
		},
		{
			name:  "another add and get",
			key:   "https://example.com/path",
			value: []byte("more test data"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(test.key, test.value)
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

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = interval + 5*time.Millisecond
	tests := []struct {
		name  string
		key   string
		value []byte
	}{
		{
			name:  "reap expired entry",
			key:   "https://example.com/expire",
			value: []byte("expired data"),
		},

		{
			name:  "unexpired entry",
			key:   "https://example.com/unexpire",
			value: []byte("unexpired data"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(test.key, test.value)
			cachedValue, found := cache.Get(test.key)
			if !found {
				t.Errorf("Expected to find key'%s' in cache", test.key)
				return
			}
			if string(cachedValue) != string(test.value) {
				t.Errorf("Expected value '%s', got '%s'", string(test.value), string(cachedValue))
				return
			}

			time.Sleep(waitTime)

			cachedValue, found = cache.Get(test.key)
			if test.name == "reap expired entry" {
				if found {
					t.Errorf("Expected key '%s' to be reaped, but it was found in cache", test.key)
				}
			}
		})
	}
}
