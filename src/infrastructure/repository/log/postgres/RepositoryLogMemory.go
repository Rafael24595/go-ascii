package log_repository

import (
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log/event"
	"go-ascii/src/infrastructure/repository"
)

const LogRepositoryLogMemory = "RepositoryLogMemory"

type RepositoryLogMemory struct {
	repository *[]log_event.LogEvent
}

func NewRepositoryLogMemory(args map[string]string) repository.RepositoryLog {
	container := dependency_container.GetInstance()
	cache := container.GetCache()
	if !cache.Exists("MEMORY_LOG") {
		cache.Put("MEMORY_LOG", "", &[]log_event.LogEvent{})
	}
	store := cache.Get("MEMORY_LOG")
	return RepositoryLogMemory{repository: store.Data().(*[]log_event.LogEvent)}
}

func (this RepositoryLogMemory) DependencyName() string {
	return LogRepositoryLogMemory
}

func (this RepositoryLogMemory) OnLoad() bool {
	return true
}

func (this RepositoryLogMemory) OnExit() bool {
	return true
}

func (this RepositoryLogMemory) FilterLog() (logs []log_event.LogEvent) {
	return *this.repository
}
