package dependency_dictionary

import (
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log/logger"
	"go-ascii/src/commons/log/logger/postgres"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/infrastructure/repository/log/postgres"
)

func FindLoggerDependency(code string, args map[string]string) logger.Logger {
	switch code {
		case logger_postgres.LoggerPostgresKey:
			return logger_postgres.NewLoggerPostgres(args)
		default:
			panic("Dependency does not exists.")
    } 
}

func FindLogDependency(code string, args map[string]string) (dependency repository.RepositoryLog) {
	switch code {
		case postgres_repository.LogRepositoryPostgresKey:
			dependency = postgres_repository.NewLogRepositoryPostgres(args)
		default:
			panic("Dependency does not exists.")
    } 
	return 
}

func FindQueryDependency(code string, args map[string]string) (dependency repository.QueryRepository) {
	switch code {
		case repository.QueryRepositoryInmemoryKey:
			dependency = repository.NewQueryRepositoryInmemory()
		default:
			panic("Dependency does not exists.")
    } 
	return 
}

func FindCommandDependency(code string, args map[string]string) (dependency repository.CommandRepository) {
	switch code {
		case repository.CommandRepositoryInmemoryKey:
			repo := dependency_container.GetInstance().GetQueryRepository()
			dependency = repository.NewCommandRepositoryInmemory(repo)
		case repository.CommandRepositoryMongoKey:
			repo := dependency_container.GetInstance().GetQueryRepository()
			dependency = repository.NewCommandRepositoryMongo(repo, args)
		default:
			panic("Dependency does not exists.")
    } 
	return 
}