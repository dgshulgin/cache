package main

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"

func Test_Cache(t *testing.T) {

	t.Run("Cache.Set",  func (t *testing.T){
		cache := NewMapCache()
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d", i)
			cache.Set(key, val)
		}
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)

			v, exists := cache.mc[key]
			if !exists {
				assert.True(t, exists)
			}
			assert.Equal(t, v, val)
		}
	})

	t.Run("Cache.Get", func (t *testing.T) {
		cache := NewMapCache()
		
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)
			cache.mc[key] = val
		}

		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)

			v, err := cache.Get(key)
			assert.NoError(t, err)
			assert.Equal(t, v, val)
		}
	})
	t.Run("Cache.Delete", func (t *testing.T) {
		cache := NewMapCache()
		
		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)
			val := fmt.Sprintf("%d",i)
			cache.mc[key] = val
		}

		for i:=0;i<10;i++{
			key := fmt.Sprintf("key#%d",i)

			cache.Delete(key)
			v, err := cache.Get(key)
			assert.Error(t, err)
			assert.Empty(t, v)
		}
	})



}

