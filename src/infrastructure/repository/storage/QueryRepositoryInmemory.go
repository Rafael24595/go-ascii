package storage

import (
	"go-ascii/src/commons/constants/log-categories"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/infrastructure/repository"
	"sort"
	"strconv"
	"sync"
)

const QueryRepositoryInmemoryKey = "QueryRepositoryInmemory"

type QueryRepositoryInmemory struct {
    sync.Mutex
    store *sync.Map
	volatile bool
}

func NewQueryRepositoryInmemory(args map[string]string) repository.QueryRepository {
	container := dependency_container.GetInstance()
	cache := container.GetCache()
	if !cache.Exists("MEMORY_REPOSITORY") {
		cache.Put("MEMORY_REPOSITORY", "", &sync.Map{})
	}
	store := cache.Get("MEMORY_REPOSITORY")
	volatile, err := strconv.ParseBool(args["GO_ASCII_QUERY_REPOSITORY_FIND_VOLATILE"])
    if err != nil {
        volatile = false
    }
	return QueryRepositoryInmemory{store: store.Data().(*sync.Map), volatile: volatile}
}

func (this QueryRepositoryInmemory) DependencyName() string {
	return QueryRepositoryInmemoryKey
}

func (this QueryRepositoryInmemory) OnLoad() bool {
	log.Log(log_categories.INFO, "Initializing \"" + this.DependencyName() + "\" dependency...")
	log.Log(log_categories.INFO, "Volatile mode for \"" + this.DependencyName() + "\" setting to " + strconv.FormatBool(this.volatile))
	log.Log(log_categories.INFO, "'" + this.DependencyName() + "' dependency was initialized successfully.")
	return true
}

func (this QueryRepositoryInmemory) OnExit() bool {
	log.Log(log_categories.INFO, "Exiting \"" + this.DependencyName() + "\" dependency...")
	log.Log(log_categories.INFO, "\"" + this.DependencyName() + "\" dependency was exited successfully.")
	return true
}

func (this QueryRepositoryInmemory) FindAll() (info []ascii.ImageInfo) {
	info = []ascii.ImageInfo{}
	this.store.Range(func(k, v interface{}) bool {
		image := v.(ascii.ImageAscii)
		data := ascii.NewImageInfo(image.GetName(), image.GetStatus(), image.GetTimestamp(), image.GetExtension())
		info = append(info, data)
		return true
	})
	sort.Slice(info[:], func(i, j int) bool {
		return info[i].GetTimestamp().UnixMilli() < info[j].GetTimestamp().UnixMilli()
	})
	return
}

func (this QueryRepositoryInmemory) Find(code string) (ascii.ImageAscii, bool) {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	if image, ok := this.store.Load(code); ok {
        cast := image.(ascii.ImageAscii)
		if this.volatile {
			this.store.Delete(code)
		}
		return cast, true
    }
    return ascii.ImageAscii{}, false
}

//@Deprecated
func (this QueryRepositoryInmemory) InsertCommand(image ascii.ImageAscii) {
}