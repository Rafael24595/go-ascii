package repository

import (
	"go-ascii/src/domain/ascii"
)

type QueryRepositoryInmemory struct {
	storage map[string]ascii.ImageAscii
}

func NewQueryRepositoryInmemory() QueryRepository {
	return QueryRepositoryInmemory{storage: map[string]ascii.ImageAscii{}}
}

func (this QueryRepositoryInmemory) FindAllAscii() (info []ascii.ImageInfo) {
	info = make([]ascii.ImageInfo, 0, len(this.storage))
	for key := range this.storage {
		image := this.storage[key]
		data := ascii.NewImageInfo(image.GetName(), image.GetExtension())
		info = append(info, data)
	}
	return
}

func (this QueryRepositoryInmemory) FindAscii(code string) ascii.ImageAscii {
	return this.storage[code]
}

func (this QueryRepositoryInmemory) InsertCommand(image ascii.ImageAscii) {
	this.storage[image.GetName()] = image
}