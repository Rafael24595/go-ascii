package dependency_dictionary

import "go-ascii/src/infrastructure/repository"

func FindQueryDependency(code string, args map[string]interface{}) repository.QueryRepository {
	switch code {
		case repository.QueryRepositoryInmemoryKey:
			return repository.NewQueryRepositoryInmemory()
		default:
			panic("Dependency does not exists.")
    } 
}

func FindCommandDependency(code string, args map[string]interface{}) repository.CommandRepository {
	switch code {
		case repository.CommandRepositoryInmemoryKey:
			repo := args["query_repository"].(repository.QueryRepository)
			return repository.NewCommandRepositoryInmemory(repo)
		case repository.CommandRepositoryMongoKey:
			repo := args["query_repository"].(repository.QueryRepository)
			return repository.NewCommandRepositoryMongo(repo)
		default:
			panic("Dependency does not exists.")
    } 
}