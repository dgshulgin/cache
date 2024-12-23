package main

import "errors"
//import "fmt"
import "sync"

func main() {
	/*
	mc := make(MapCache, 2)
	mc.Set("first","1")
	v, err := mc.Get("first")
	fmt.Printf("key=%s, val=%s\n", "first", v)
	
	mc.Set("second", "2")
	v, err = mc.Get("second")
	fmt.Printf("key=%s, val=%s\n", "second", v)

	mc.Delete("first")
	v, err = mc.Get("first")
	fmt.Printf("err=%v\n", err)
	*/
}

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


