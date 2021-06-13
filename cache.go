package lcache

import (
	"errors"
	"sync"
	"time"
)

var cacheRegistry *sync.Map

var errCacheNotExist = errors.New("local cache :key not found")

func init() {
	cacheRegistry = &sync.Map{}
	cacheCleanup() //init clean routine
}

func cacheCleanup() {
	go func() {
		cleanExpire()
	}()
}

func cleanExpire() {
	cacheRegistry.Range(func(key, val interface{}) bool {
		cache := val.(*LocalCache)
		cache.clean()
		return true
	})
}

//cache
type LocalCache struct {
	cache *sync.Map
}

type LocalCacheEntry struct {
	data   interface{}
	expire int64 //expire =0 永不过期
}

func NewLocalCache() *LocalCache {
	return &LocalCache{
		cache: &sync.Map{},
	}
}

func (c *LocalCache) Get(key string) (interface{}, error) {
	val, ok := c.cache.Load(key)
	if !ok {
		return nil, errCacheNotExist
	}
	entry := val.(*LocalCacheEntry)
	if entry.expire == 0 {
		return entry.data, nil
	}
	if time.Now().Unix() > entry.expire {
		c.cache.Delete(key)
		return nil, errCacheNotExist
	}
	return entry.data, nil
}

func (c *LocalCache) Set(key string, val interface{}, expire int64) error {
	if key == "" {
		return errors.New("local cache: key is empty ")
	}
	if val == nil {
		return errors.New("local cache: val is nil ")
	}
	if expire < 0 {
		expire = 0
	} else {
		expire = time.Now().Unix() + expire
	}
	c.cache.Store(key, &LocalCacheEntry{val, expire})
	return nil
}

func (c *LocalCache) clean() {
	c.cache.Range(func(key, val interface{}) bool {
		entry := val.(*LocalCacheEntry)
		if entry.expire == 0 {
			return true
		}
		if time.Now().Unix() > entry.expire {
			c.cache.Delete(key)
		}
		return true
	})
}
