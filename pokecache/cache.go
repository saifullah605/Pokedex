package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	slots map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       any
}

func NewCache(interval time.Duration) Cache {

	c := Cache{map[string]cacheEntry{}, &sync.Mutex{}}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{time.Now(), val}
	c.slots[key] = entry
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.slots[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()

		for key, val := range c.slots {
			if now.Sub(val.createdAt) > interval {
				delete(c.slots, key)
			}
		}
		c.mu.Unlock()
	}

}
