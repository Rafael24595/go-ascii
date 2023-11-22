package logger_repository

import (
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log/event"
	"go-ascii/src/commons/log/logger"

	_ "github.com/lib/pq"
)

const LoggerMemoryKey = "LoggerMemory"

type LoggerMemory struct {
	repository *[]log_event.LogEvent
}

func NewLoggerMemory(args map[string]string) logger.Logger {
	container := dependency_container.GetInstance()
	cache := container.GetCache()
	if !cache.Exists("MEMORY_LOG") {
		cache.Put("MEMORY_LOG", "", &[]log_event.LogEvent{})
	}
	store := cache.Get("MEMORY_LOG")
	return &LoggerMemory{repository: store.Data().(*[]log_event.LogEvent)}
}

func (this LoggerMemory) DependencyName() string {
	return LoggerMemoryKey
}

func (this LoggerMemory) OnLoad() bool {
	return true
}

func (this LoggerMemory) OnExit() bool {
	return true
}

func (this *LoggerMemory) Log(event log_event.LogEvent) string {
	*this.repository = append(*this.repository, event)
	return ""
}