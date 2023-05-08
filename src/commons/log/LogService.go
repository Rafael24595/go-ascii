package log

import (
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/constants/log-categories"
	"go-ascii/src/commons/log/event"
	"go-ascii/src/commons/log/logger"
	"time"
)

type LogService struct {
	logger logger.Logger
}

var logService *LogService

func Log(category log_categories.LogCategory, message string) string {
	if logService == nil {
		panic("Not instanced")
	}
	configuration := configuration.GetInstance()
	sessionId := configuration.GetSessionId()
	date := time.Now()
	event := log_event.NewLogEvent(sessionId, string(category), "SYSTEM", message, date)
	return logService.logger.Log(event)
}

func LogFam(category log_categories.LogCategory, family string, message string) string {
	if logService == nil {
		panic("Not instanced")
	}
	configuration := configuration.GetInstance()
	sessionId := configuration.GetSessionId()
	date := time.Now()
	event := log_event.NewLogEvent(sessionId, string(category), family, message, date)
	return logService.logger.Log(event)
}

func Instance(logger logger.Logger) *LogService {
	if logService == nil {
		logService = &LogService{logger: logger}
		return logService
	}
	panic("Already instanced")
}