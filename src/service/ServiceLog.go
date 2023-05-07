package service

import (
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/repository"
)

type ServiceLog struct {
	logRepository repository.RepositoryLog
}

func NewServiceLog(logRepository repository.RepositoryLog) ServiceLog {
	return ServiceLog{logRepository: logRepository}
}

func (this ServiceLog) FindAll() (response []dto.InfoLogResponse) {
	logs := this.logRepository.FindAll()
	response = []dto.InfoLogResponse{}
	for _, log := range logs {
		response = append(response, dto.InfoLogResponse{SessionId: log.GetSessionId(),Category: log.GetCategory(), Message: log.GetMessage(), Timestamp: log.GetTimestamp()})
	}
	return
}

func (this ServiceLog) Find(category string) (response []dto.InfoLogResponse) {
	logs := this.logRepository.Find(category)
	response = []dto.InfoLogResponse{}
	for _, log := range logs {
		response = append(response, dto.InfoLogResponse{SessionId: log.GetSessionId(),Category: log.GetCategory(), Message: log.GetMessage(), Timestamp: log.GetTimestamp()})
	}
	return
}