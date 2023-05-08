package service

import (
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/commons/constants/request-state"
)

type ServiceAscii struct {
	queryRepository repository.QueryRepository
	commandRepository repository.CommandRepository
	requestLauncher RequestLauncher
}

func NewServiceAscii(queryRepository repository.QueryRepository, commandRepository repository.CommandRepository) ServiceAscii {
	requestLauncher := NewRequestLauncher(commandRepository);
	return ServiceAscii{queryRepository: queryRepository, commandRepository: commandRepository, requestLauncher: requestLauncher}
}

func (this ServiceAscii) FindAll() (response []dto.InfoResponse) {
	info := this.queryRepository.FindAll()
	response = []dto.InfoResponse{}
	for _, data := range info {
		response = append(response, dto.InfoResponse{Code: data.GetCode(), Status: data.GetStatus(), Extension: data.GetExtension()})
	}
	return
}

func (this ServiceAscii) Find(code string) dto.InfoAsciiResponse {
	image := this.queryRepository.Find(code)
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

func (this ServiceAscii) Insert(dto dto.ImageRequest) string {
	return this.requestLauncher.PushAsciiRequest(dto)
}

func (this ServiceAscii) Modify(code string) string {
	image := this.queryRepository.Find(code)
	image.SetStatus(request_state.RESTORED)
	return this.commandRepository.Modify(image)
}

func (this ServiceAscii) Delete(code string) string {
	image := this.queryRepository.Find(code)
	image.SetStatus(request_state.DELETED)
	return this.commandRepository.Delete(image)
}