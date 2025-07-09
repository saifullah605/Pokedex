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
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	c := Cache{map[string]cacheEntry{}, &sync.Mutex{}}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{time.Now(), val}
	c.slots[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.slots[key]
	return entry.val, ok
}

func (c *Cache) reapLoop() {

}
