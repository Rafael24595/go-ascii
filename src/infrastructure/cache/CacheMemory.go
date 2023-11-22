package cache

const CacheMemoryKey = "CacheMemory"

type CacheMemory struct {
    cache map[string]*CacheEvent[interface{}]
}

func NewCacheMemory(args map[string]string) Cache {
    return &CacheMemory{cache: map[string]*CacheEvent[interface{}]{}}
}

func (self CacheMemory) DependencyName() string {
	return CacheMemoryKey
}

func (self CacheMemory) OnLoad() bool {
	return true
}

func (self CacheMemory) OnExit() bool {
	return true
}

func (self CacheMemory) Exists(key string) bool {
    data := self.cache[key]
    return data != nil
}

func (self CacheMemory) Get(key string) *CacheEvent[interface{}] {
    return self.cache[key]
}

func (self *CacheMemory) Put(key string, reference string, value interface{}) *CacheEvent[interface{}] {
    event := NewCacheEvent(key, reference, value)
    self.cache[key] = &event;
    return &event;
}

func (self *CacheMemory) Delete(key string) *CacheEvent[interface{}] {
    value := self.Get(key)
    delete(self.cache, key)
    return value
}