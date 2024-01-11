package cache

import (
	"context"

	"github.com/ant-go/framework/contracts/icache"
	"github.com/ant-go/framework/contracts/icodec"
	"github.com/spf13/cast"
)

type Cache[T any] struct {
	codec icodec.CodecInterface
	store icache.StoreInterface
}

func New[T any](codec icodec.CodecInterface, store icache.StoreInterface) *Cache[T] {
	return &Cache[T]{
		codec: codec,
		store: store,
	}
}

func (c *Cache[T]) Get(ctx context.Context, key any) (res T, exists bool, err error) {
	cacheKey := c.getCacheKey(key)

	var value any
	value, exists, err = c.store.Get(ctx, cacheKey)
	if err != nil {
		return
	}

	if c.store.Codec() {
		err = c.codec.Unmarshal(nil, &res)
		exists = err == nil
	} else {
		if v, ok := value.(T); ok {
			res, exists = v, true
			return
		}
	}

	return
}

func (c *Cache[T]) Set(ctx context.Context, key any, value T, options ...icache.StoreOption) (err error) {
	cacheKey := c.getCacheKey(key)

	if c.store.Codec() {
		var data []byte
		if data, err = c.codec.Marshal(&value); err == nil {
			err = c.store.Set(ctx, cacheKey, data, options...)
		}
	} else {
		err = c.store.Set(ctx, cacheKey, value, options...)
	}

	return
}

func (c *Cache[T]) Delete(ctx context.Context, key any) (err error) {
	cacheKey := c.getCacheKey(key)

	err = c.store.Delete(ctx, cacheKey)

	return
}

func (c *Cache[T]) getCacheKey(key any) string {
	switch v := key.(type) {
	case icache.CacheKey:
		return v.GetCacheKey()
	default:
		return cast.ToString(key)
	}
}
