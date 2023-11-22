package cache

import "go-ascii/src/commons"

type Cache interface {
	commons.Dependency
	Exists(key string) bool
	Get(key string) *CacheEvent[interface{}]
	Put(key string, reference string, value interface{}) *CacheEvent[interface{}]
	Delete(key string) *CacheEvent[interface{}]
}