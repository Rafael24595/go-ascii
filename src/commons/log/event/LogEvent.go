package log_event

import "time"

type LogEvent struct {
	registerId int
	sessionId string
	category string
	message string
	timestamp time.Time
}

func NewLogEvent(sessionId string, category string, message string, timestamp time.Time) LogEvent {
	return LogEvent{sessionId: sessionId, category: category, message: message, timestamp: timestamp}
}

func NewLogEventDB(registerId int, sessionId string, category string, message string, timestamp time.Time) LogEvent {
	return LogEvent{registerId: registerId, sessionId: sessionId, category: category, message: message, timestamp: timestamp}
}

func (this LogEvent) GetId() int {
	return this.registerId
}

func (this LogEvent) GetSessionId() string {
	return this.sessionId
}

func (this LogEvent) GetCategory() string {
	return this.category
}

func (this LogEvent) GetMessage() string {
	return this.message
}

func (this LogEvent) GetTimestamp() time.Time {
	return this.timestamp
}