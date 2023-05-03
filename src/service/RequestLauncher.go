package service

import (
	"path/filepath"
	"golang.org/x/exp/slices"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/commons/utils"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/repository"
)

type RequestLauncher struct {
	commandRepository repository.CommandRepository
	pending *[]ascii.QueueEvent
	process *[]ascii.QueueEvent
	failed *[]ascii.QueueEvent
	success *[]string
	active * bool
}

func NewRequestLauncher(commandRepository repository.CommandRepository) RequestLauncher {
	pending := []ascii.QueueEvent{}
	process := []ascii.QueueEvent{}
	success := []string{}
	failed := []ascii.QueueEvent{}
	active := false
	return RequestLauncher{commandRepository: commandRepository, pending: &pending, process: &process, success: &success, failed: &failed, active: &active}
}

func (this RequestLauncher) CheckStatus(code string) (string , string) {
	var index int
	index = slices.IndexFunc(*this.pending, func(e ascii.QueueEvent) bool { return e.GetCode() == code})
	if index != -1 {
		return request_state.PENDING, ""
	}
	index = slices.IndexFunc(*this.process, func(e ascii.QueueEvent) bool { return e.GetCode() == code})
	if index != -1 {
		return request_state.PROCESS, ""
	}
	index = slices.IndexFunc(*this.failed, func(e ascii.QueueEvent) bool { return e.GetCode() == code})
	if index != -1 {
		return request_state.FAILED, (*this.failed)[index].GetMessage()
	}
	index = slices.IndexFunc(*this.success, func(c string) bool { return c == code})
	if index != -1 {
		return request_state.SUCCES, ""
	}
	return request_state.NOT_FOUND, ""
}

func (this RequestLauncher) PushAsciiRequest(dto dto.ImageRequest) string {
	path := tempsource.Base64ToSource(dto.Image, dto.Code)
	event := ascii.NewQueueEvent(dto, path)
	*this.pending = append(*this.pending, event)
	go this.launchQueuque()
	return filepath.Base(path)
}

func (this *RequestLauncher) launchQueuque() {
	if !*this.active {
		*this.active = true
		for *this.active {
			pend := (*this.pending)[0]
			go this.insertAscii(pend)
			*this.pending = (*this.pending)[1:]
			*this.process = append(*this.process, pend)
			if len(*this.pending) == 0 {
				*this.active = false
			}
		}
	}
}

func (this RequestLauncher) insertAscii(event ascii.QueueEvent) {
	builderAscii, err := builder.NewBuilderAscii(event)
	if err != nil {
		event.SetMessage(err.Error())
		this.markAsFailed(event)
	} else {
		imageAscii := builderAscii.Build()
		this.commandRepository.InsertAscii(imageAscii)
		this.markAsComplete(event)
	}
}

func (this RequestLauncher) markAsComplete(event ascii.QueueEvent) {
	idx := slices.IndexFunc(*this.process, func(e ascii.QueueEvent) bool { return e.GetCode() == event.GetCode()})
	if idx != -1 {
		*this.process = utils.RemoveIndex(*this.process, idx)
	}

	*this.success = append(*this.success, event.GetCode())
}

func (this RequestLauncher) markAsFailed(event ascii.QueueEvent) {
	idx := slices.IndexFunc(*this.process, func(e ascii.QueueEvent) bool { return e.GetCode() == event.GetCode()})
	if idx != -1 {
		*this.process = utils.RemoveIndex(*this.process, idx)
	}

	*this.failed = append(*this.failed, event)
}