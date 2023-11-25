package dependency_dictionary

import (
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log/logger"
	"go-ascii/src/commons/log/logger/postgres"
	"go-ascii/src/infrastructure/cache"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/infrastructure/repository/log/postgres"
	"go-ascii/src/infrastructure/repository/storage"
)

func FindLoggerDependency(code string, args map[string]string) logger.Logger {
	switch code {
		case logger_repository.LoggerPostgresKey:
			return logger_repository.NewLoggerPostgres(args)
		case logger_repository.LoggerMemoryKey:
			return logger_repository.NewLoggerMemory(args)
		default:
			panic("Dependency does not exists.")
    } 
}

func FindLogDependency(code string, args map[string]string) (dependency repository.RepositoryLog) {
	switch code {
		case log_repository.LogRepositoryPostgresKey:
			dependency = log_repository.NewLogRepositoryPostgres(args)
		case log_repository.LogRepositoryLogMemory:
			dependency = log_repository.NewRepositoryLogMemory(args)
		default:
			panic("Dependency does not exists.")
    } 
	return 
}

func FindQueryDependency(code string, args map[string]string) (dependency repository.QueryRepository) {
	switch code {
		case storage.QueryRepositoryInmemoryKey:
			dependency = storage.NewQueryRepositoryInmemory(args)
		default:
			panic("Dependency does not exists.")
    } 
	return 
}

func FindCommandDependency(code string, args map[string]string) (dependency repository.CommandRepository) {
	switch code {
		case storage.CommandRepositoryInmemoryKey:
			dependency = storage.NewCommandRepositoryInmemory(args)
		case storage.CommandRepositoryMongoKey:
			repo := dependency_container.GetInstance().GetQueryRepository()
			dependency = storage.NewCommandRepositoryMongo(repo, args)
		default:
			panic("Dependency does not exists.")
    } 
	return 
}

func FindCacheDependency(code string, args map[string]string) (dependency cache.Cache) {
	switch code {
		case cache.CacheMemoryKey:
			return cache.NewCacheMemory(args)
		default:
			panic("Dependency does not exists.")
    }
}