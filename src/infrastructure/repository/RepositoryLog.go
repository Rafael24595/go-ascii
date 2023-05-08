package repository

import (
	"go-ascii/src/commons"
	"go-ascii/src/commons/log/event"
)

type RepositoryLog interface {
	commons.Dependency
	FilterLog() []log_event.LogEvent
}