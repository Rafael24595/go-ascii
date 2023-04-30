package repository

import "go-ascii/src/domain/ascii"

type RepositoryInmemory struct {
	storage map[string]ascii.ImageAscii
}

func NewRepositoryInmemory() Repository {
	return RepositoryInmemory{storage: map[string]ascii.ImageAscii{}}
}

func (this RepositoryInmemory) FindAllAscii() (result string) {
	return "HelloWorld!"
}

func (this RepositoryInmemory) FindAscii(code string) ascii.ImageAscii {
	return this.storage[code]
}

func (this RepositoryInmemory) InsertAscii(image ascii.ImageAscii) string {
	this.storage[image.Name] = image
	return image.Name
}