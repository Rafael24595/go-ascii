package repository

import (
	"sort"
	"go-ascii/src/domain/ascii"
)

const QueryRepositoryInmemoryKey = "QueryRepositoryInmemory"

type QueryRepositoryInmemory struct {
	storage map[string]ascii.ImageAscii
}

func NewQueryRepositoryInmemory() QueryRepository {
	return QueryRepositoryInmemory{storage: map[string]ascii.ImageAscii{}}
}

func (this QueryRepositoryInmemory) DependencyName() string {
	return QueryRepositoryInmemoryKey
}

func (this QueryRepositoryInmemory) OnLoad() bool {
	return true
}

func (this QueryRepositoryInmemory) OnExit() bool {
	return true
}

func (this QueryRepositoryInmemory) FindAll() (info []ascii.ImageInfo) {
	info = make([]ascii.ImageInfo, 0, len(this.storage))
	for key := range this.storage {
		image := this.storage[key]
		data := ascii.NewImageInfo(image.GetName(), image.GetStatus(), image.GetTimestamp(), image.GetExtension())
		info = append(info, data)
	}
	sort.Slice(info[:], func(i, j int) bool {
		return info[i].GetTimestamp().UnixMicro() < info[j].GetTimestamp().UnixMicro()
	})
	return
}

func (this QueryRepositoryInmemory) Find(code string) ascii.ImageAscii {
	return this.storage[code]
}

func (this QueryRepositoryInmemory) InsertCommand(image ascii.ImageAscii) {
	this.storage[image.GetName()] = image
}