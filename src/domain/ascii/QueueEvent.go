package ascii

import (
	"path/filepath"
	"go-ascii/src/infrastructure/dto"
)

type QueueEvent struct {
	path  string
	code  string
	image string
	message string
}

func NewQueueEvent(dto dto.ImageRequest, path string) QueueEvent {
	return QueueEvent{path: path, code: filepath.Base(path), image: dto.Image, message: ""}
}

func (this QueueEvent) GetPath() string {
	return this.path
}

func (this QueueEvent) GetCode() string {
	return this.code
}

func (this QueueEvent) GetImage() string {
	return this.image
}

func (this QueueEvent) GetMessage() string {
	return this.message
}

func (this *QueueEvent) SetMessage(message string) {
	this.message = message
}