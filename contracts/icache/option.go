package icache

import (
	"time"
)

type StoreOption func(*StoreOptions)

type StoreOptions struct {
	Expiration time.Duration
	Cost       int64
}

func ApplyStoreOptions(options ...StoreOption) *StoreOptions {
	o := new(StoreOptions)
	for _, fn := range options {
		fn(o)
	}
	return o
}

func WithStoreExpiration(expiration time.Duration) StoreOption {
	return func(options *StoreOptions) {
		options.Expiration = expiration
	}
}

func WithStoreCost(cost int64) StoreOption {
	return func(options *StoreOptions) {
		options.Cost = cost
	}
}
