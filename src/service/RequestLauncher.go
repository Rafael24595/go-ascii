package service

import (
	"os"
	"path/filepath"
	"golang.org/x/exp/slices"
	"go-ascii/src/commons/constants"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/commons/utils"
	"go-ascii/src/commons/utils/image"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder"
	"go-ascii/src/infrastructure/dto"
	"go-ascii/src/infrastructure/repository"
)

type RequestLauncher struct {
	commandRepository repository.CommandRepository
	pending *[]ascii.QueueEvent
	process *[]ascii.QueueEvent
	failed *[]ascii.QueueEvent
	success *[]string
	active bool
}

func NewRequestLauncher(commandRepository repository.CommandRepository) RequestLauncher {
	pending := []ascii.QueueEvent{}
	process := []ascii.QueueEvent{}
	success := []string{}
	failed := []ascii.QueueEvent{}
	return RequestLauncher{commandRepository: commandRepository, pending: &pending, process: &process, success: &success, failed: &failed, active: false}
}

func (this RequestLauncher) CheckStatus(code string) string {
	idx := slices.IndexFunc(*this.pending, func(e ascii.QueueEvent) bool { return e.GetCode() == code})
	if idx != -1 {
		return constants.PENDING
	}
	idx = slices.IndexFunc(*this.process, func(e ascii.QueueEvent) bool { return e.GetCode() == code})
	if idx != -1 {
		return constants.PENDING
	}
	idx = slices.IndexFunc(*this.failed, func(e ascii.QueueEvent) bool { return e.GetCode() == code})
	if idx != -1 {
		return constants.FAILED
	}
	idx = slices.IndexFunc(*this.success, func(c string) bool { return c == code})
	if idx != -1 {
		return constants.SUCCES
	}
	
	return constants.NOT_FOUND
}

func (this RequestLauncher) PushAsciiRequest(dto dto.AsciiRequest) string {
	path := tempsource.Base64ToSource(dto.Image, dto.Code)
	event := ascii.NewQueueEvent(dto, path)
	*this.pending = append(*this.pending, event)
	go this.launchQueuque()
	return filepath.Base(path)
}

func (this RequestLauncher) launchQueuque() {
	if !this.active {
		this.active = true
		for this.active {
			pend := (*this.pending)[0]
			go this.insertAscii(pend)
			*this.pending = utils.RemoveIndex(*this.pending, 0)
			*this.process = append(*this.process, pend)
			if len(*this.pending) == 0 {
				this.active = false
			}
		}
	}
}

func (this RequestLauncher) insertAscii(event ascii.QueueEvent) {
	temp, err := os.Open(event.GetPath())
	if err != nil {
		panic(err)
	}

	img := image.Decode(temp)
	scaleHeight := 115
	scaleWidth := 0
	grayScale := constants.GrayScaleLevels["default"]

	builderAscii := builder.NewBuilderAscii(img, scaleHeight, scaleWidth, grayScale)
	imageAscii := builderAscii.Build()

	temp, err = os.Open(event.GetPath())
	if err != nil {
		panic(err)
	}

	imageAscii.Name = event.GetCode()
	imageAscii.Type = utils.FileExtension(temp)

	temp.Close()

	this.commandRepository.InsertAscii(imageAscii)
	
	idx := slices.IndexFunc(*this.process, func(e ascii.QueueEvent) bool { return e.GetCode() == event.GetCode()})
	if idx != -1 {
		*this.process = utils.RemoveIndex(*this.process, idx)
	}

	*this.success = append(*this.success, event.GetCode())
}