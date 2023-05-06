package configurator

import (
	"os"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/dependency-container/dependency-dictionary"
)

func LoadConfiguration() (Configuration, dependency_container.DependencyContainer) {
	configuration := buildConfiguration()
	dependencyContainer := buildDependencyContainer()
	return configuration, dependencyContainer
}

func buildConfiguration() Configuration {
	configuration := GetInstance()
	configuration.domain = os.Getenv("GO_ASCII_DOMAIN")
	configuration.port = os.Getenv("GO_ASCII_PORT")
	return *configuration
}

func buildDependencyContainer() dependency_container.DependencyContainer {
	dependencyContainer := dependency_container.GetInstance()

	queryRepositoryKey := os.Getenv("GO_ASCII_QUERY_REPOSITORY")
	queryRepositoryArgs := map[string]interface{}{}
	queryRepository := dependency_dictionary.FindQueryDependency(queryRepositoryKey, queryRepositoryArgs)

	commandRepositoryKey := os.Getenv("GO_ASCII_COMMAND_REPOSITORY")
	commandRepositoryArgs := map[string]interface{}{"query_repository": queryRepository}
	commandRepository := dependency_dictionary.FindCommandDependency(commandRepositoryKey, commandRepositoryArgs)

	dependencyContainer.SetQueryRepository(queryRepository)
	dependencyContainer.SetCommandRepository(commandRepository)

	dependencyContainer.OnLoad()

	return *dependencyContainer
}