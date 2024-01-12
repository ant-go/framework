package redis

import (
	"context"
	"errors"
	"time"

	"github.com/ant-go/framework/contracts"
	"github.com/ant-go/framework/contracts/icache"
	"github.com/ant-go/framework/contracts/ierror"
	"github.com/ant-go/framework/implements/cache/store"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
)

type Client interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

type Store struct {
	client Client
}

func New(client Client) *Store {
	return &Store{
		client: client,
	}
}

func (s *Store) Type() string {
	return store.TypeRedis
}

func (s *Store) Codec() bool {
	return true
}

func (s *Store) Get(ctx context.Context, key any) (value any, exists bool, err error) {
	var stringKey string
	if stringKey, err = validateKey(key); err != nil {
		return
	}

	res := s.client.Get(ctx, stringKey)
	if errors.Is(res.Err(), redis.Nil) {
		err = ierror.New(icache.ErrNotFound, "key", key)
		return
	}

	if err = res.Err(); err != nil {
		err = ierror.New(icache.ErrGetFailed, "err", err)
		return
	}

	value, exists = res.Val(), true

	return
}

func (s *Store) Set(ctx context.Context, key, value any, options ...icache.StoreOption) (err error) {
	var o *icache.StoreOptions
	contracts.ApplyOptions(o, options)

	var stringKey string
	if stringKey, err = validateKey(key); err != nil {
		return
	}

	err = s.client.Set(ctx, stringKey, value, o.Expiration).Err()
	if err != nil {
		err = ierror.New(icache.ErrSetFailed, "key", key, "value", value, "ttl", o.Expiration)
	}

	return
}

func (s *Store) Delete(ctx context.Context, key any) (err error) {
	var stringKey string
	if stringKey, err = validateKey(key); err != nil {
		return
	}

	err = s.client.Del(ctx, stringKey).Err()
	if err != nil {
		err = ierror.New(icache.ErrDelFailed, "key", key)
	}

	return
}

func validateKey(key any) (stringKey string, err error) {
	if stringKey = cast.ToString(key); stringKey == "" {
		err = ierror.New(icache.ErrEmptyKey)
	}
	return
}
