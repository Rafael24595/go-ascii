package logger

import (
	"go-ascii/src/commons"
	"go-ascii/src/commons/log/event"
)

type Logger interface {
	commons.Dependency
	Log(event log_event.LogEvent) string
}