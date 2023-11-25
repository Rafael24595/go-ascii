package storage

import (
	"go-ascii/src/commons/constants/log-categories"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/infrastructure/repository"
	"strconv"
	"sync"
)

const CommandRepositoryInmemoryKey = "CommandRepositoryInmemory"

type CommandRepositoryInmemory struct {
	sync.Mutex
    store *sync.Map
	stateless bool
}

func NewCommandRepositoryInmemory(args map[string]string) repository.CommandRepository {
	container := dependency_container.GetInstance()
	cache := container.GetCache()
	if !cache.Exists("MEMORY_REPOSITORY") {
		cache.Put("MEMORY_REPOSITORY", "", &sync.Map{})
	}
	store := cache.Get("MEMORY_REPOSITORY")
	stateless, err := strconv.ParseBool(args["GO_ASCII_COMMAND_REPOSITORY_STATELESS"])
    if err != nil {
        stateless = false
    }
	return &CommandRepositoryInmemory{store: store.Data().(*sync.Map), stateless: stateless}
}

func (this *CommandRepositoryInmemory) DependencyName() string {
	return CommandRepositoryInmemoryKey
}

func (this *CommandRepositoryInmemory) OnLoad() bool {
	log.Log(log_categories.INFO, "Initializing \"" + this.DependencyName() + "\" dependency...")
	log.Log(log_categories.INFO, "Stateless mode for \"" + this.DependencyName() + "\" setting to " + strconv.FormatBool(this.stateless))
	log.Log(log_categories.INFO, "\"" + this.DependencyName() + "\" dependency was initialized successfully.")
	return true
}

func (this *CommandRepositoryInmemory) OnExit() bool {
	log.Log(log_categories.INFO, "Exiting \"" + this.DependencyName() + "\" dependency...")
	log.Log(log_categories.INFO, "\"" + this.DependencyName() + "\" dependency was exited successfully.")
	return true
}

func (this *CommandRepositoryInmemory) Insert(image ascii.ImageAscii) string {
	this.store.Store(image.GetName(), image)
	return image.GetName()
}

func (this *CommandRepositoryInmemory) Modify(image ascii.ImageAscii) string {
	image.SetStatus(request_state.RESTORED)
	this.store.Store(image.GetName(), image)
	return image.GetName()
}

func (this *CommandRepositoryInmemory) Delete(image ascii.ImageAscii) string {
	if this.stateless {
		this.store.Delete(image.GetName())
	} else {
		image.SetStatus(request_state.DELETED)
		this.store.Store(image.GetName(), image)
	}
	return image.GetName()
}

func (this *CommandRepositoryInmemory) ToQuery(image ascii.ImageAscii) {
	//TODO: Remove implementation
}