package dependency_container

import (
	"go-ascii/src/infrastructure/repository"
)

type DependencyContainer struct {
	logRepository repository.RepositoryLog
	queryRepository repository.QueryRepository
	commandRepository repository.CommandRepository
}

var dependencyContainer *DependencyContainer

func GetInstance() *DependencyContainer {
	if dependencyContainer == nil {
		dependencyContainer = &DependencyContainer{}
	}
	return dependencyContainer
}

func (this *DependencyContainer) GetLogRepository() repository.RepositoryLog {
	if this.logRepository != nil {
		return this.logRepository
	}
	panic("Dependency not found.")
}

func (this *DependencyContainer) SetLogRepository(dependency repository.RepositoryLog ) {
	this.logRepository = dependency
}

func (this *DependencyContainer) GetQueryRepository() repository.QueryRepository {
	if this.queryRepository != nil {
		return this.queryRepository
	}
	panic("Dependency not found.")
}

func (this *DependencyContainer) SetQueryRepository(dependency repository.QueryRepository ) {
	this.queryRepository = dependency
}

func (this *DependencyContainer) GetCommandRepository() repository.CommandRepository {
	if this.commandRepository != nil {
		return this.commandRepository
	}
	panic("Dependency not found.")
}

func (this *DependencyContainer) SetCommandRepository(dependency repository.CommandRepository ) {
	this.commandRepository = dependency
}

func (this *DependencyContainer) OnLoad() {
	this.queryRepository.OnLoad()
	this.commandRepository.OnLoad()
}

func (this *DependencyContainer) OnExit() {
	this.queryRepository.OnExit()
	this.commandRepository.OnExit()
}