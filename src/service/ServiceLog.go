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

func (this ServiceLog) FilterLog(logParams dto.LogParamsRequest) (response []dto.InfoLogResponse) {
	logFilter := this.buildLogFilter(logParams)
	logs := this.logRepository.FilterLog()
	logs = logFilter.Filter(logs)
	response = []dto.InfoLogResponse{}
	for _, log := range logs {
		response = append(response, dto.NewInfoLogResponse(log.GetId(), log.GetSessionId(), log.GetCategory(), log.GetFamily(), log.GetMessage(), log.GetTimestamp().UnixMilli()))
	}
	return
}

func (this ServiceLog) buildLogFilter(logParams dto.LogParamsRequest) log.LogFilter {
	category := logParams.Category
	family := logParams.Family

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

	return log.NewLogFilter(category, family, from, to)
}