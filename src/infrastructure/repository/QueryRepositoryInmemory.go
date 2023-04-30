package repository

import "go-ascii/src/domain/ascii"

type QueryRepositoryInmemory struct {
	storage map[string]ascii.ImageAscii
}

func NewQueryRepositoryInmemory() QueryRepository {
	return QueryRepositoryInmemory{storage: map[string]ascii.ImageAscii{}}
}

func (this QueryRepositoryInmemory) FindAllAscii() (keys []string) {
	keys = make([]string, 0, len(this.storage))
	for k := range this.storage {
		keys = append(keys, k)
	}
	return
}

func (this QueryRepositoryInmemory) FindAscii(code string) ascii.ImageAscii {
	return this.storage[code]
}

func (this QueryRepositoryInmemory) InsertCommand(image ascii.ImageAscii) {
	this.storage[image.Name] = image
}