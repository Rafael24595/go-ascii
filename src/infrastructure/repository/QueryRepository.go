package repository

import (
	"go-ascii/src/commons"
	"go-ascii/src/domain/ascii"
)

type QueryRepository interface {
	commons.Dependency
	FindAllAscii() []ascii.ImageInfo
	FindAscii(code string) ascii.ImageAscii
	InsertCommand(ascii.ImageAscii)
}