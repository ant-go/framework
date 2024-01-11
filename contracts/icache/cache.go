package icache

import (
	"context"
)

type CacheKey interface {
	GetCacheKey() string
}

type CacheInterface[T any] interface {
	Get(ctx context.Context, key any) (res T, exists bool, err error)
	Set(ctx context.Context, key any, value T, options ...StoreOption) (err error)
	Delete(ctx context.Context, key any) (err error)
}

type StoreInterface interface {
	Type() string
	Codec() bool
	Get(ctx context.Context, key any) (value any, exists bool, err error)
	Set(ctx context.Context, key, value any, options ...StoreOption) (err error)
	Delete(ctx context.Context, key any) (err error)
}
