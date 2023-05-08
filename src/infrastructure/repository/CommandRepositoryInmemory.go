package repository

import (
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/log"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/commons/constants/log-categories"
)

const CommandRepositoryInmemoryKey = "CommandRepositoryInmemory"

type CommandRepositoryInmemory struct {
	queryRepository QueryRepository
}

func NewCommandRepositoryInmemory(queryRepository QueryRepository) CommandRepository {
	return CommandRepositoryInmemory{queryRepository: queryRepository}
}

func (this CommandRepositoryInmemory) DependencyName() string {
	return CommandRepositoryInmemoryKey
}

func (this CommandRepositoryInmemory) OnLoad() bool {
	log.Log(log_categories.INFO, "Initializing \"" + this.DependencyName() + "\" dependency...")
	log.Log(log_categories.INFO, "\"" + this.DependencyName() + "\" dependency was initialized successfully.")
	return true
}

func (this CommandRepositoryInmemory) OnExit() bool {
	log.Log(log_categories.INFO, "Exiting \"" + this.DependencyName() + "\" dependency...")
	log.Log(log_categories.INFO, "\"" + this.DependencyName() + "\" dependency was exited successfully.")
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