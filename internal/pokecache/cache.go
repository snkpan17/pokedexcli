package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	rwmutex  sync.RWMutex
	done     chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
		done:     make(chan struct{}),
	}
	newCache.reapLoop()
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.rwmutex.Lock()
	defer c.rwmutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.rwmutex.RLock()
	defer c.rwmutex.RUnlock()
	if entry, ok := c.cache[key]; ok {
		return entry.val, ok
	}
	return nil, false
}

func cleanMap(c *Cache) {
	c.rwmutex.Lock()
	defer c.rwmutex.Unlock()
	for key, entry := range c.cache {
		since := time.Since(entry.createdAt)
		if since > c.interval {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop() {
	timer := time.NewTicker(c.interval)
	go func() {
		for {
			select {
			case <-timer.C:
				cleanMap(c)
			case <-c.done:
				timer.Stop()
				return
			}

		}
	}()
}

func (c *Cache) Stop() {
	close(c.done)
}
