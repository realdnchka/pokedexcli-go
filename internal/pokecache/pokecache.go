package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	cacheMap map[string]cacheEntry
	mu sync.RWMutex 
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *cache {
	c := cache {
		cacheMap: map[string]cacheEntry{},
		mu: sync.RWMutex{},
	}

	go reapLoop(&c, interval)
	return &c
}

func (c *cache) Add(key string, v []byte) {
	c.mu.RLock()
	c.cacheMap[key] = cacheEntry { createdAt: time.Now(), val: v }
	c.mu.RUnlock()
}

func (c *cache) Get(key string) ([]byte, bool) {
	if _, ok := c.cacheMap[key]; !ok {
		return nil, false
	}

	if c.cacheMap[key].val == nil {
		return nil, false
	}

	return c.cacheMap[key].val, true
}

func reapLoop(c *cache, interval time.Duration) {
	
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for k, v := range c.cacheMap {
			if time.Since(v.createdAt) > interval {
				c.cacheMap[k] = cacheEntry{}
			}
		}
	}
}
