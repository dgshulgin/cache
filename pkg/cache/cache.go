package cache

import "errors"
import "sync"
import "context"

var (
    ErrNotFound = errors.New("value not found")
    ErrTimeout = errors.New("operation timeout")
    ErrOperationFailed = errors.New("operation failed")
)

type Cache interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}


type MapCacheStorage map[string]string
type MapCache struct {
	mc MapCacheStorage
	mut sync.Mutex
}

func NewMapCache() Cache { //MapCache {
	return &MapCache{
		mc: make(MapCacheStorage),
	}
}

func (c *MapCache) Set(ctx context.Context, key, value string) error {
    ch := make(chan struct{})
    go func(){
        //set
        defer close(ch)
        c.mut.Lock()
        c.mc[key] = value
        c.mut.Unlock()
        ch <- struct{}{}
    }()
    //check timeout or set
    select {
        case <-ctx.Done():
            return ErrTimeout
        case _, ok := <-ch:
            if ok {
                return nil
            }
            return ErrOperationFailed
    }
}

func (c MapCache) Get(ctx context.Context, key string) (string, error) {
    ch := make(chan string)
    go func(){
        defer close(ch)

        c.mut.Lock()
        defer c.mut.Unlock()

        if _, exists := c.mc[key]; exists {
            ch <- c.mc[key]
        }
    }()

    select {
       case <-ctx.Done():
           return "", ErrTimeout
       case v, ok := <- ch:
            if ok {
                return v, nil
            }
            return "", ErrNotFound
    }
}

func (c MapCache) Delete(ctx context.Context, key string) error {
    ch := make(chan struct{})
    go func(){
        defer close(ch)
        
        c.mut.Lock()
        defer c.mut.Unlock()

        if _, exists := c.mc[key]; exists {
            delete(c.mc, key)
        }
        ch <- struct{}{}
    }()

    select {
        case <-ctx.Done():
            return ErrTimeout
        case _, ok := <-ch:
            if ok {
                return nil
            }
            return ErrNotFound
    }
}


