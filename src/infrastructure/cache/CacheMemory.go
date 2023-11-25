package cache

import "sync"

const CacheMemoryKey = "CacheMemory"

type CacheMemory struct {
    sync.Mutex
    cache sync.Map
}

func NewCacheMemory(args map[string]string) Cache {
	return &CacheMemory{cache: sync.Map{}}
}

func (self *CacheMemory) DependencyName() string {
	return CacheMemoryKey
}

func (self *CacheMemory) OnLoad() bool {
	return true
}

func (self *CacheMemory) OnExit() bool {
	return true
}

func (self *CacheMemory) Exists(key string) bool {
     _, ok := self.cache.Load(key)
    return ok
}

func (self *CacheMemory) Get(key string) *CacheEvent[interface{}] {
    if self.Exists(key) {
        event, _ := self.cache.Load(key)
        return event.(*CacheEvent[interface{}])
    }
    return nil 
}

func (self *CacheMemory) Put(key string, reference string, value interface{}) *CacheEvent[interface{}] {
	event := NewCacheEvent(key, reference, value)
    self.cache.Store(key, &event)
	return &event
}

func (self *CacheMemory) Delete(key string) *CacheEvent[interface{}] {
	value := self.Get(key)
    self.cache.Delete(key)
	return value
}