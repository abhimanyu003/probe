package cache

import (
	"time"

	"github.com/muesli/cache2go"
)

type Cache struct {
	cache *cache2go.CacheTable
}

func NewCache(table string) *Cache {
	return &Cache{
		cache: cache2go.Cache(table),
	}
}

func (c *Cache) Set(key string, value any) {
	c.cache.Add(key, time.Hour*1, value)
}

func (c *Cache) Get(key string) (any, error) {
	item, err := c.cache.Value(key)
	if err != nil {
		return nil, err
	}

	return item.Data(), nil
}

func (c *Cache) Flush() {
	c.cache.Flush()
}
