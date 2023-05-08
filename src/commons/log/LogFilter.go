package log

import (
	"go-ascii/src/commons/log/event"
	"strings"
	"time"
)

type LogFilter struct {
	category string
	from time.Time
	to time.Time
}

func NewLogFilter(category string, from time.Time, to time.Time) LogFilter {
	return LogFilter{category: category, from: from, to: to}
}

func (this LogFilter) Filter(events []log_event.LogEvent) (filterEvents []log_event.LogEvent) {
	filterEvents = events
	filterEvents = this.categoryFilter(filterEvents)
	filterEvents = this.fromFilter(filterEvents)
	filterEvents = this.toFilter(filterEvents)
	return
}

func (this LogFilter) categoryFilter(events []log_event.LogEvent) (filterEvents []log_event.LogEvent) {
	if this.category == "" {
		return events
	}

	filterEvents = []log_event.LogEvent{}
	for _, event := range events {
		if  strings.EqualFold(event.GetCategory(), this.category) {
			filterEvents = append(filterEvents, event)
		}
	}
	return 
}

func (this LogFilter) fromFilter(events []log_event.LogEvent) (filterEvents []log_event.LogEvent) {
	if this.from.Equal(time.Time{}) {
		return events
	}

	fromMs := this.from.UnixMilli()
	filterEvents = []log_event.LogEvent{}
	for _, event := range events {
		eventMs := event.GetTimestamp().UnixMilli()
		if eventMs >= fromMs {
			filterEvents = append(filterEvents, event)
		}
	}
	return 
}

func (this LogFilter) toFilter(events []log_event.LogEvent) (filterEvents []log_event.LogEvent) {
	if this.to.Equal(time.Time{}) {
		return events
	}

	toMs := this.to.UnixMilli()
	filterEvents = []log_event.LogEvent{}
	for _, event := range events {
		eventMs := event.GetTimestamp().UnixMilli()
		if eventMs <= toMs {
			filterEvents = append(filterEvents, event)
		}
	}
	return 
}