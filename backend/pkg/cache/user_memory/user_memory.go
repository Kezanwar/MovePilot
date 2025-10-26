package user_memory_cache

import (
	user_repo "movepilot/pkg/repositories/user"
	"sync"
	"time"
)

type cacheItem struct {
	user      *user_repo.Model
	expiresAt time.Time
}

type Cache struct {
	store map[string]cacheItem
	mutex sync.RWMutex
	ttl   time.Duration
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		store: make(map[string]cacheItem),
		ttl:   ttl,
	}
}

// Get a user from cache, returns nil if not found or expired
func (c *Cache) Get(uuid string) *user_repo.Model {
	c.mutex.RLock()
	item, found := c.store[uuid]
	c.mutex.RUnlock()

	if !found {
		return nil
	}

	if time.Now().After(item.expiresAt) {
		// remove expired item
		c.mutex.Lock()
		delete(c.store, uuid)
		c.mutex.Unlock()
		return nil
	}

	return item.user
}

// Set stores a user in the cache
func (c *Cache) Set(uuid string, u *user_repo.Model) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.store[uuid] = cacheItem{
		user:      u,
		expiresAt: time.Now().Add(c.ttl),
	}
}

// Delete removes a user from the cache
func (c *Cache) Delete(uuid string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.store, uuid)
}
