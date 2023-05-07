package configuration

import (
	"strconv"
	"time"
)

const serviceName = "go-ascii"

var configuration *Configuration

type Configuration struct {
	args map[string]string
	serviceName string
	sesionId string
	timestamp time.Time
	domain string
	port string
}

func GetInstance() *Configuration {
	if configuration == nil {
		panic("Not instanced")
	}
	return configuration
}

func Instance(args map[string]string) *Configuration {
	if configuration == nil {
		timestamp := time.Now()
		miliseconds := timestamp.UnixNano() / int64(time.Millisecond)
		configuration = &Configuration{}
		configuration.serviceName = serviceName
		configuration.sesionId = serviceName + "-" + strconv.FormatInt(miliseconds, 10)
		configuration.timestamp = timestamp
		configuration.args = args
		configuration.domain = args["GO_ASCII_DOMAIN"]
		configuration.port = args["GO_ASCII_PORT"]
		return configuration
	}
	panic("Already instanced")
}

func (this Configuration) GetServiceName() string {
	return this.serviceName
}

func (this Configuration) GetSessionId() string {
	return this.sesionId
}

func (this Configuration) GetTimestamp() time.Time {
	return this.timestamp
}

func (this Configuration) GetAddr() string {
	return this.domain + ":" + this.port
}

func (this Configuration) GetDomain() string {
	return this.domain
}

func (this Configuration) GetPort() string {
	return this.port
}

func (this Configuration) GetArg(key string) string {
	return this.args[key]
}

func (this Configuration) GetArgs() map[string]string {
	return this.args
}