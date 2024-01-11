package ristretto

import (
	"context"
	"time"

	"github.com/ant-go/framework/contracts/icache"
	"github.com/ant-go/framework/contracts/ierror"
	"github.com/ant-go/framework/implements/cache/store"
	_ "github.com/dgraph-io/ristretto"
)

type Client interface {
	Get(key any) (any, bool)
	SetWithTTL(key, value any, cost int64, ttl time.Duration) bool
	Del(key any)
	Wait()
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
	return store.TypeRistretto
}

func (s *Store) Codec() bool {
	return false
}

func (s *Store) Get(_ context.Context, key any) (value any, exists bool, err error) {
	value, exists = s.client.Get(key)

	if !exists {
		err = ierror.New(icache.ErrNotFound, "key", key)
	}

	return
}

func (s *Store) Set(_ context.Context, key, value any, options ...icache.StoreOption) (err error) {
	var o icache.StoreOptions
	for _, option := range options {
		option(&o)
	}

	if ok := s.client.SetWithTTL(key, value, o.Cost, o.Expiration); !ok {
		err = ierror.New(icache.ErrSetFailed, "key", key, "value", value, "ttl", o.Expiration, "cost", o.Cost)
	}

	if err != nil {
		return
	}

	s.client.Wait()

	return
}

func (s *Store) Delete(_ context.Context, key any) (err error) {
	s.client.Del(key)
	return
}
