package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]chacheEntry
	mux   *sync.Mutex
}

type chacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache: make(map[string]chacheEntry),
		mux:   &sync.Mutex{},
	}
	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = chacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	value, exists := c.cache[key]
	return value.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
