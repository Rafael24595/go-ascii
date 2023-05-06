package configurator

var configuration *Configuration

type Configuration struct {
	domain string
	port string
}

func GetInstance() *Configuration {
	if configuration == nil {
		configuration = &Configuration{}
	}
	return configuration
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