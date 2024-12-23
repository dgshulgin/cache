package cache

import "testing"
import "fmt"
//import "github.com/stretchr/testify/assert"
import "sync"
import "context"
import "time"

func Benchmark_Cache_Set(b *testing.B) {

    timeout := time.Millisecond
    ctx,_ := context.WithTimeout(context.Background(), timeout)

    cache := NewMapCache()
	for i:=0;i<b.N;i++ {
		key := fmt.Sprintf("key#%d",i)
		val := fmt.Sprintf("%d",i)
		cache.Set(ctx, key, val)
	}

}

func Benchmark_CacheConcurrent(b *testing.B) {
    
    timeout := time.Millisecond
    ctx,_ := context.WithTimeout(context.Background(), timeout)

    wg:= sync.WaitGroup{}
    cache := NewMapCache()
  for i:=0; i< b.N;i++ {
	    key := fmt.Sprintf("key#%d",i)
	    val := fmt.Sprintf("%d",i)
	
    	wg.Add(1)    
    	go func() {
		    cache.Set(ctx, key,val)
    		wg.Done()
    	}()
  }

  wg.Wait()
} 
