package dependency_container

import (
	"go-ascii/src/commons"
	"go-ascii/src/commons/constants/log-categories"
	"go-ascii/src/commons/log"
	"go-ascii/src/infrastructure/cache"
	"go-ascii/src/infrastructure/repository"
)

type DependencyContainer struct {
	logRepository repository.RepositoryLog
	queryRepository repository.QueryRepository
	commandRepository repository.CommandRepository
	cache cache.Cache
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
	this.logLoadDependency(dependency)
}

func (this *DependencyContainer) GetQueryRepository() repository.QueryRepository {
	if this.queryRepository != nil {
		return this.queryRepository
	}
	panic("Dependency not found.")
}

func (this *DependencyContainer) SetQueryRepository(dependency repository.QueryRepository ) {
	this.queryRepository = dependency
	this.logLoadDependency(dependency)
}

func (this *DependencyContainer) GetCommandRepository() repository.CommandRepository {
	if this.commandRepository != nil {
		return this.commandRepository
	}
	panic("Dependency not found.")
}

func (this *DependencyContainer) SetCommandRepository(dependency repository.CommandRepository ) {
	this.commandRepository = dependency
	this.logLoadDependency(dependency)
}

func (this *DependencyContainer) GetCache() cache.Cache {
	if this.cache != nil {
		return this.cache
	}
	panic("Dependency not found.")
}

func (this *DependencyContainer) SetCache(dependency cache.Cache) {
	this.cache = dependency
}

func (this *DependencyContainer) OnLoad() {
	log.Log(log_categories.INFO, "Initializing dependencies...")
	this.queryRepository.OnLoad()
	this.commandRepository.OnLoad()
	log.Log(log_categories.INFO, "Dependencies initialized successfully.")
}

func (this *DependencyContainer) OnExit() {
	this.queryRepository.OnExit()
	this.commandRepository.OnExit()
}

func (this *DependencyContainer) logLoadDependency(dependency commons.Dependency) {
	log.Log(log_categories.INFO, "Dependency \"" + dependency.DependencyName() + "\" loaded.")
}