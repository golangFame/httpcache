package inmem

import (
	memcache "github.com/bxcodec/gotcha/cache"
	"github.com/bxcodec/hache/cache"
)

type inmemCache struct {
	cache memcache.Cache
}

// NewCache ...
func NewCache(c memcache.Cache) cache.Interactor {
	return &inmemCache{
		cache: c,
	}
}

func (i *inmemCache) Set(key string, value cache.CachedResponse) (err error) {
	return i.cache.Set(key, value)
}

func (i *inmemCache) Get(key string) (res cache.CachedResponse, err error) {
	item, err := i.cache.Get(key)
	if err != nil {
		return
	}
	res = item.(cache.CachedResponse)
	return
}

func (i *inmemCache) Delete(key string) (err error) {
	return i.cache.Delete(key)
}