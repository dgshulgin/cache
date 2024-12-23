package cache

import "testing"
import "fmt"
//import "github.com/stretchr/testify/assert"
import "sync"

func Benchmark_Cache_Set(b *testing.B) {
	cache := NewMapCache()
	for i:=0;i<b.N;i++ {
		key := fmt.Sprintf("key#%d",i)
		val := fmt.Sprintf("%d",i)
		cache.Set(key, val)
	}

}

func Benchmark_CacheConcurrent(b *testing.B) {
	wg:= sync.WaitGroup{}
	cache := NewMapCache()
  for i:=0; i< b.N;i++ {
	

	key := fmt.Sprintf("key#%d",i)
	val := fmt.Sprintf("%d",i)
	
	wg.Add(1)    
	go func() {
		cache.Set(key,val)
		wg.Done()
	}()

  }

  wg.Wait()
} 
