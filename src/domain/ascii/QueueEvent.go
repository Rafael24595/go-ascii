package ascii

import (
	"path/filepath"
	"go-ascii/src/infrastructure/dto"
)

type QueueEvent struct {
	path  string
	code  string
	image string
}

func NewQueueEvent(dto dto.AsciiRequest, path string) QueueEvent {
	return QueueEvent{path: path, code: filepath.Base(path), image: dto.Image}
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