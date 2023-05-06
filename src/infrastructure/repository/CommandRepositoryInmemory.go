package repository

import (
	request_state "go-ascii/src/commons/constants/request-state"
	"go-ascii/src/domain/ascii"
)

const CommandRepositoryInmemoryKey = "CommandRepositoryInmemory"

type CommandRepositoryInmemory struct {
	queryRepository QueryRepository
}

func NewCommandRepositoryInmemory(queryRepository QueryRepository) CommandRepository {
	return CommandRepositoryInmemory{queryRepository: queryRepository}
}

func (this CommandRepositoryInmemory) OnLoad() bool {
	return true
}

func (this CommandRepositoryInmemory) OnExit() bool {
	return true
}

func (this CommandRepositoryInmemory) Insert(image ascii.ImageAscii) string {
	this.ToQuery(image)
	return image.GetName()
}

func (this CommandRepositoryInmemory) Modify(image ascii.ImageAscii) string {
	image.SetStatus(request_state.RESTORED)
	this.ToQuery(image)
	return image.GetName()
}

func (this CommandRepositoryInmemory) Delete(image ascii.ImageAscii) string {
	image.SetStatus(request_state.DELETED)
	this.ToQuery(image)
	return image.GetName()
}

func (this CommandRepositoryInmemory) ToQuery(image ascii.ImageAscii) {
	this.queryRepository.InsertCommand(image)
}