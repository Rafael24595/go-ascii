package configurator

import (
	"os"
	"strings"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/dependency-container/dependency-dictionary"
)

func LoadConfiguration() (Configuration, dependency_container.DependencyContainer) {
	rawConfig := loadArgsFromEnv()
	configuration := buildConfiguration(rawConfig)
	dependencyContainer := buildDependencyContainer(rawConfig)
	return configuration, dependencyContainer
}

func buildConfiguration(rawConfig map[string]string) Configuration {
	configuration := GetInstance()
	configuration.args = rawConfig
	configuration.domain = rawConfig["GO_ASCII_DOMAIN"]
	configuration.port = rawConfig["GO_ASCII_PORT"]
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

func buildDependencyContainer(rawConfig map[string]string) dependency_container.DependencyContainer {
	configuration := GetInstance()
	args := configuration.GetArgs()
	dependencyContainer := dependency_container.GetInstance()

	queryRepositoryKey := configuration.GetArg("GO_ASCII_QUERY_REPOSITORY")
	queryRepository := dependency_dictionary.FindQueryDependency(queryRepositoryKey, args)
	dependencyContainer.SetQueryRepository(queryRepository)

	commandRepositoryKey := configuration.GetArg("GO_ASCII_COMMAND_REPOSITORY")
	commandRepository := dependency_dictionary.FindCommandDependency(commandRepositoryKey, args)
	dependencyContainer.SetCommandRepository(commandRepository)


	dependencyContainer.OnLoad()

	return *dependencyContainer
}