package cache

import "time"

type CacheEvent[T any] struct {
	key       string
	reference string
	timestamp int
	data      T
}

func NewCacheEvent[T any](key string, reference string, data T) CacheEvent[T] {
	return CacheEvent[T]{key: key, reference: reference, timestamp: int(time.Now().UnixMilli()), data: data}
}

func (self CacheEvent[T]) Key() string {
	return self.key
}

func (self CacheEvent[T]) Reference() string {
	return self.reference
}

func (self CacheEvent[T]) Timestamp() int {
	return self.timestamp
}

func (self CacheEvent[T]) Data() T {
	return self.data
}