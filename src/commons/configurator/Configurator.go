package configurator

import (
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/constants/log-categories"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/dependency-container/dependency-dictionary"
	"go-ascii/src/commons/log"
	"go-ascii/src/infrastructure/cache"
	"os"
	"strings"
)

func LoadConfiguration() (configuration.Configuration, dependency_container.DependencyContainer) {
	rawConfig := loadArgsFromEnv()
	configuration := buildConfiguration(rawConfig)
	_ = buildCacheDependency(rawConfig)
	_ = buildLogDependency(rawConfig)
	log.Log(log_categories.INFO, "Loading configuration...")
	log.Log(log_categories.INFO, "Session id established: "+ configuration.GetSessionId() +"")
	log.Log(log_categories.INFO, "Start date: "+ configuration.GetTimestamp().String() +"")
	dependencyContainer := buildDependencyContainer(rawConfig)
	log.Log(log_categories.INFO, "Configuration loaded successfully.")
	return configuration, dependencyContainer
}

func buildConfiguration(rawConfig map[string]string) configuration.Configuration {
	configuration := configuration.Instance(rawConfig)
	return *configuration
}

func loadArgsFromEnv() map[string]string {
	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
        items := make(map[string]string)
        for _, item := range data {
            key, val := getkeyval(item)
            items[key] = val
        }
        return items
    }
    return getenvironment(os.Environ(), func(item string) (key, val string) {
        splits := strings.Split(item, "=")
        key = splits[0]
        val = splits[1]
        return
    })
}

func buildCacheDependency(rawConfig map[string]string) *cache.Cache {
	configuration := configuration.GetInstance()
	args := configuration.GetArgs()

	cacheKey := configuration.GetArg("GO_ASCII_CACHE")
	cacheDependency := dependency_dictionary.FindCacheDependency(cacheKey, args)
	cacheDependency.OnLoad()

	dependencyContainer := dependency_container.GetInstance()
	dependencyContainer.SetCache(cacheDependency)

	return &cacheDependency
}

func buildLogDependency(rawConfig map[string]string) log.LogService {
	configuration := configuration.GetInstance()
	args := configuration.GetArgs()

	loggerKey := configuration.GetArg("GO_ASCII_LOGGER")
	loggerDependency := dependency_dictionary.FindLoggerDependency(loggerKey, args)
	service := log.Instance(loggerDependency)
	loggerDependency.OnLoad()

	return *service
}

func buildDependencyContainer(rawConfig map[string]string) dependency_container.DependencyContainer {
	configuration := configuration.GetInstance()
	args := configuration.GetArgs()
	dependencyContainer := dependency_container.GetInstance()

	logRepositoryKey := configuration.GetArg("GO_ASCII_LOG_REPOSITORY")
	logRepositoryDependency := dependency_dictionary.FindLogDependency(logRepositoryKey, args)
	dependencyContainer.SetLogRepository(logRepositoryDependency)
	
	queryRepositoryKey := configuration.GetArg("GO_ASCII_QUERY_REPOSITORY")
	queryRepositoryDependency := dependency_dictionary.FindQueryDependency(queryRepositoryKey, args)
	dependencyContainer.SetQueryRepository(queryRepositoryDependency)
	
	commandRepositoryKey := configuration.GetArg("GO_ASCII_COMMAND_REPOSITORY")
	commandRepositoryDependency := dependency_dictionary.FindCommandDependency(commandRepositoryKey, args)
	dependencyContainer.SetCommandRepository(commandRepositoryDependency)
	
	dependencyContainer.OnLoad()

	return *dependencyContainer
}