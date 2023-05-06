package dependency_dictionary

import (
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/infrastructure/repository"
)

func FindQueryDependency(code string) repository.QueryRepository {
	switch code {
		case repository.QueryRepositoryInmemoryKey:
			return repository.NewQueryRepositoryInmemory()
		default:
			panic("Dependency does not exists.")
    } 
}

func FindCommandDependency(code string) repository.CommandRepository {
	switch code {
		case repository.CommandRepositoryInmemoryKey:
			repo := dependency_container.GetInstance().GetQueryRepository()
			return repository.NewCommandRepositoryInmemory(repo)
		case repository.CommandRepositoryMongoKey:
			repo := dependency_container.GetInstance().GetQueryRepository()
			return repository.NewCommandRepositoryMongo(repo)
		default:
			panic("Dependency does not exists.")
    } 
}