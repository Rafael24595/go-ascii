package service

import (
	"go-ascii/src/commons/dto"
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

func (this Service) FindAllAscii() (response []dto.InfoResponse) {
	info := this.queryRepository.FindAllAscii()
	response = []dto.InfoResponse{}
	for _, data := range info {
		response = append(response, dto.InfoResponse{Code: data.GetCode(), Extension: data.GetExtension()})
	}
	return
}

func (this Service) FindAscii(code string) dto.AsciiResponse {
	image := this.queryRepository.FindAscii(code)
	response := dto.AsciiResponse{Name: image.GetName(), Extension: image.GetExtension(), Status: image.GetStatus(), Frames: image.GetFrames()}
	if response.Status == "" {
		status, message := this.requestLauncher.CheckStatus(code)
		response.Status = status
		response.Message = message
		response.Name = code
		response.Extension = "Undefined"
	}
	return response
}

func (this Service) InsertAscii(dto dto.ImageRequest) string {
	return this.requestLauncher.PushAsciiRequest(dto)
}