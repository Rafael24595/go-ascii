package service

import (
	"go-ascii/src/domain/ascii"
	"go-ascii/src/infrastructure/dto"
	"go-ascii/src/infrastructure/repository"
)

type Service struct {
	queryRepository repository.QueryRepository
	requestLauncher RequestLauncher
}

func NewService(queryRepository repository.QueryRepository, commandRepository repository.CommandRepository) Service {
	requestLauncher := NewRequestLauncher(commandRepository);
	return Service{queryRepository: queryRepository, requestLauncher: requestLauncher}
}

func (this Service) FindAllAscii() []ascii.ImageInfo {
	return this.queryRepository.FindAllAscii()
}

func (this Service) FindAscii(code string) ascii.ImageAscii {
	image := this.queryRepository.FindAscii(code)
	if image.Status == "" {
		status := this.requestLauncher.CheckStatus(code)
		image.Status = status
		image.Name = code
	}
	return image
}

func (this Service) InsertAscii(dto dto.AsciiRequest) string {
	return this.requestLauncher.PushAsciiRequest(dto)
}