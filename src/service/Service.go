package service

import (
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/repository"
)

type Service struct {
	queryRepository repository.QueryRepository
	commandRepository repository.CommandRepository
	requestLauncher RequestLauncher
}

func NewService(queryRepository repository.QueryRepository, commandRepository repository.CommandRepository) Service {
	requestLauncher := NewRequestLauncher(commandRepository);
	return Service{queryRepository: queryRepository, commandRepository: commandRepository, requestLauncher: requestLauncher}
}

func (this Service) FindAll() (response []dto.InfoResponse) {
	info := this.queryRepository.FindAllAscii()
	response = []dto.InfoResponse{}
	for _, data := range info {
		response = append(response, dto.InfoResponse{Code: data.GetCode(), Status: data.GetStatus(), Extension: data.GetExtension()})
	}
	return
}

func (this Service) Find(code string) dto.InfoAsciiResponse {
	image := this.queryRepository.FindAscii(code)
	height, width := image.GetDimensions()
	response := dto.InfoAsciiResponse{Name: image.GetName(), Height: height, Width: width, Extension: image.GetExtension(), Status: image.GetStatus(), Frames: image.GetFrames()}
	if response.Status == "" {
		status, message := this.requestLauncher.CheckStatus(code)
		response.Status = status
		response.Message = message
		response.Name = code
		response.Extension = "Undefined"
	}
	return response
}

func (this Service) Insert(dto dto.ImageRequest) string {
	return this.requestLauncher.PushAsciiRequest(dto)
}

func (this Service) Modify(code string) string {
	image := this.queryRepository.FindAscii(code)
	image.SetStatus(request_state.RESTORED)
	return this.commandRepository.Modify(image)
}

func (this Service) Delete(code string) string {
	image := this.queryRepository.FindAscii(code)
	image.SetStatus(request_state.DELETED)
	return this.commandRepository.Delete(image)
}