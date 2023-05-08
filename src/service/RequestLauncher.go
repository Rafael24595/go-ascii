package service

import (
	"path/filepath"
	"golang.org/x/exp/slices"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/log"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/commons/utils"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/commons/constants/log-categories"
)

type RequestLauncher struct {
	commandRepository repository.CommandRepository
	pending *[]ascii.QueueEvent
	process *[]ascii.QueueEvent
	failed *[]ascii.QueueEvent
	success *[]string
	active * bool
}

const family = "ASCII-QUEUE"

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
	log.LogFam(log_categories.INFO, family, "Event with code \"" + event.GetCode() + "\" added to pending queue.")
	go this.launchQueuque()
	return filepath.Base(path)
}

func (this *RequestLauncher) launchQueuque() {
	if !*this.active {
		*this.active = true
		log.LogFam(log_categories.INFO, family, "Running queue.")
		for *this.active {
			pend := (*this.pending)[0]
			log.LogFam(log_categories.INFO, family, "Launching routine for \"" + pend.GetCode() + "\".")
			go this.insertAscii(pend)
			*this.pending = (*this.pending)[1:]
			*this.process = append(*this.process, pend)
			if len(*this.pending) == 0 {
				*this.active = false
				log.LogFam(log_categories.INFO, family, "Queue is empty.")
			}
		}
		log.LogFam(log_categories.INFO, family, "Queue stopped.")
	}
}

func (this RequestLauncher) insertAscii(event ascii.QueueEvent) {
	builderAscii, err := builder.NewBuilderAscii(event)
	if err != nil {
		event.SetMessage(err.Error())
		this.markAsFailed(event)
	} else {
		imageAscii := builderAscii.Build()
		this.commandRepository.Insert(imageAscii)
		this.markAsComplete(event)
	}
	log.LogFam(log_categories.INFO, family, "Routine for \"" + event.GetCode() + "\" ended.")
}

func (this RequestLauncher) markAsComplete(event ascii.QueueEvent) {
	idx := slices.IndexFunc(*this.process, func(e ascii.QueueEvent) bool { return e.GetCode() == event.GetCode()})
	if idx != -1 {
		*this.process = utils.RemoveIndex(*this.process, idx)
	}

	*this.success = append(*this.success, event.GetCode())
	tempsource.CleanSource(event.GetCode())
	log.LogFam(log_categories.INFO, family, "File \"" +  event.GetCode() + "\" transformed into ASCII successfully.")
}

func (this RequestLauncher) markAsFailed(event ascii.QueueEvent) {
	idx := slices.IndexFunc(*this.process, func(e ascii.QueueEvent) bool { return e.GetCode() == event.GetCode()})
	if idx != -1 {
		*this.process = utils.RemoveIndex(*this.process, idx)
	}

	*this.failed = append(*this.failed, event)
	log.LogFam(log_categories.ERROR, family, "Cannot transform into ASCII file \"" +  event.GetCode() + "\". Context: " + event.GetMessage())
}