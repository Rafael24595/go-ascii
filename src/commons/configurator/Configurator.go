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
	queryRepository := dependency_dictionary.FindQueryDependency(queryRepositoryKey)
	dependencyContainer.SetQueryRepository(queryRepository)

	commandRepositoryKey := os.Getenv("GO_ASCII_COMMAND_REPOSITORY")
	commandRepository := dependency_dictionary.FindCommandDependency(commandRepositoryKey)
	dependencyContainer.SetCommandRepository(commandRepository)


	dependencyContainer.OnLoad()

	return *dependencyContainer
}