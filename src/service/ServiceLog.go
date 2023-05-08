package service

import (
	"time"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/log"
	"go-ascii/src/commons/utils"
	"go-ascii/src/infrastructure/repository"
)

type ServiceLog struct {
	logRepository repository.RepositoryLog
}

func NewServiceLog(logRepository repository.RepositoryLog) ServiceLog {
	return ServiceLog{logRepository: logRepository}
}

func (this ServiceLog) FindAll(logParams dto.LogParamsRequest) (response []dto.InfoLogResponse) {
	logFilter := this.buildLogFilter(logParams)
	logs := this.logRepository.FindAll()
	logs = logFilter.Filter(logs)
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

func (this ServiceLog) buildLogFilter(logParams dto.LogParamsRequest) log.LogFilter {
	category := logParams.Category

	fromMilis, err := utils.ParseInt64(logParams.From)
    if err != nil {
        panic(err)
    }
	from := time.Time{}
	if fromMilis != 0 {
		from = time.Unix(0, int64(fromMilis) * int64(time.Millisecond))
	}

	toMilis, err := utils.ParseInt64(logParams.To)
    if err != nil {
        panic(err)
    }
	to := time.Time{}
	if toMilis != 0 {
		to = time.Unix(0, int64(toMilis) * int64(time.Millisecond))
	}

	return log.NewLogFilter(category, from, to)
}