package cache

import (
	"errors"
	"time"
)

type Cache struct {
	items map[string]Item
}

type Item struct {
	Value   interface{}
	Created time.Time
}

func New() *Cache {
	items := make(map[string]Item)

	return &Cache{
		items: items,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) error {
	if _, ok := c.items[key]; !ok {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}
