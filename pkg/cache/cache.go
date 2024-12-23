package cache

import "errors"
import "sync"

var ErrNotFound = errors.New("value not found")

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}


type MapCacheStorage map[string]string
type MapCache struct {
	mc MapCacheStorage
	mut sync.Mutex
}

func NewMapCache() MapCache {
	return MapCache{
		mc: make(MapCacheStorage),
	}
}


func (c *MapCache) Set(key string, value string) error {
	c.mut.Lock()
	c.mc[key] = value
	c.mut.Unlock()
	return nil
}

func (c MapCache) Get(key string) (string, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	if _, exists := c.mc[key]; exists {
		return c.mc[key], nil
	}
	return "", ErrNotFound
}

func (c *MapCache) Delete(key string) error {
	c.mut.Lock()
	defer c.mut.Unlock()

	if _, exists := c.mc[key]; exists {
		delete(c.mc, key)
	}
	return ErrNotFound
}


