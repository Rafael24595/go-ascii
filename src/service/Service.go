package service

import "go-ascii/src/infrastructure/repository"

type Service struct {
	Repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return Service{Repository: repository}
}

func (this Service) FindAllAscii() (result string) {
	return this.Repository.FindAllAscii()
}