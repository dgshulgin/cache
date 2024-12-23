package cache

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"
//import "github.com/dgshulgin/cache/pkg/cache"
import "context"
import "time"


func Test_Cache(t *testing.T) {

	t.Run("Cache.Set",  func (t *testing.T){

        timeout := time.Millisecond

       baseC := context.Background()
       ctx,_ := context.WithTimeout(baseC, timeout)

        cache := NewMapCache()
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d", i)
			cache.Set(ctx, key, val)
		}
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)

			v, exists := cache.(*MapCache).mc[key]
			if !exists {
				assert.True(t, exists)
			}
			assert.Equal(t, v, val)
		}
	})

	t.Run("Cache.Get", func (t *testing.T) {
        timeout := time.Millisecond

       baseC := context.Background()
       ctx,_ := context.WithTimeout(baseC, timeout)
		cache := NewMapCache()
		
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)
			cache.(*MapCache).mc[key] = val
		}

		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)

			v, err := cache.Get(ctx, key)
			assert.NoError(t, err)
			assert.Equal(t, v, val)
		}
	})
	t.Run("Cache.Delete", func (t *testing.T) {
        timeout := time.Millisecond

       baseC := context.Background()
       ctx,_ := context.WithTimeout(baseC, timeout)
		cache := NewMapCache()
		
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)
			cache.(*MapCache).mc[key] = val
		}

		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)

			cache.Delete(ctx, key)
			v, err := cache.Get(ctx, key)
			assert.Error(t, err)
			assert.Empty(t, v)
		}
	})



}

