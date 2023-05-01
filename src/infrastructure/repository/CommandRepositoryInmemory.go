package repository

import "go-ascii/src/domain/ascii"

type CommandRepositoryInmemory struct {
	queryRepository QueryRepository
}

func NewCommandRepositoryInmemory(queryRepository QueryRepository) CommandRepository {
	return CommandRepositoryInmemory{queryRepository: queryRepository}
}

func (this CommandRepositoryInmemory) InsertAscii(image ascii.ImageAscii) string {
	this.InsertQuery(image)
	return image.GetName()
}

func (this CommandRepositoryInmemory) InsertQuery(image ascii.ImageAscii) {
	this.queryRepository.InsertCommand(image)
}