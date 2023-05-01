package repository

import "go-ascii/src/domain/ascii"

type QueryRepository interface {
	FindAllAscii() []ascii.ImageInfo
	FindAscii(code string) ascii.ImageAscii
	InsertCommand(ascii.ImageAscii)
}