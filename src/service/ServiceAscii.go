package service

import (
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/infrastructure/repository"
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

func (this ServiceAscii) Find(code string) (dto.InfoAsciiResponse, bool) {
	image, ok := this.queryRepository.Find(code)
	if !ok {
		image = ascii.ImageAscii{}
	}
	height, width := image.GetDimensions()
	response := dto.InfoAsciiResponse{Name: image.GetName(), Height: height, Width: width, Extension: image.GetExtension(), Status: image.GetStatus(), Frames: image.GetFrames()}
	if response.Status == "" {
		status, message := this.requestLauncher.CheckStatus(code)
		response.Status = status
		response.Message = message
		response.Name = code
		response.Extension = "Undefined"
	}
	return response, ok
}

func (this ServiceAscii) Insert(dto dto.ImageRequest) string {
	return this.requestLauncher.PushAsciiRequest(dto)
}

func (this ServiceAscii) Modify(code string) (string, bool) {
	image, ok := this.queryRepository.Find(code)
	if !ok {
		return code, false
	}
	image.SetStatus(request_state.RESTORED)
	return this.commandRepository.Modify(image), true
}

func (this ServiceAscii) Delete(code string) (string, bool) {
	image, ok := this.queryRepository.Find(code)
	if !ok {
		return code, false
	}
	image.SetStatus(request_state.DELETED)
	return this.commandRepository.Delete(image), true
}