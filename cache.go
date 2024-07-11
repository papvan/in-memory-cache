package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	items           map[string]Item
	cleanupInterval time.Duration
}

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

func New(cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)

	cache := Cache{
		items:           items,
		cleanupInterval: cleanupInterval,
	}

	if cleanupInterval > 0 {
		cache.startGC()
	}

	return &cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	var expirianced int64

	if ttl > 0 {
		expirianced = time.Now().Add(ttl).UnixNano()
	}

	c.items[key] = Item{
		Value:      value,
		Expiration: expirianced,
	}
	c.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.items[key]; !ok {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}

func (c *Cache) startGC() {
	go c.gc()
}

func (c *Cache) gc() {
	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		// Ищем элементы с истекшим временем жизни и удаляем из хранилища
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}

	}
}

// expiredKeys returns expired keys
func (c *Cache) expiredKeys() (keys []string) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

// clearItems removes item that are expired by keys
func (c *Cache) clearItems(keys []string) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
