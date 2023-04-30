package repository

type RepositoryInmemory struct {
}

func NewRepositoryInmemory() Repository {
	return RepositoryInmemory{}
}

func (this RepositoryInmemory) FindAllAscii() (result string) {
	return "HelloWorld!"
}