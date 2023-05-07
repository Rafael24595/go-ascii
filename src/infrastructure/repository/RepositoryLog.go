package repository

import (
	"go-ascii/src/commons"
	"go-ascii/src/commons/log/event"
)

type RepositoryLog interface {
	commons.Dependency
	FindAll() []log_event.LogEvent
	Find(cattegory string) []log_event.LogEvent
}