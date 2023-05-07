package repository

import (
	"go-ascii/src/commons"
	"go-ascii/src/domain/ascii"
)

type QueryRepository interface {
	commons.Dependency
	FindAll() []ascii.ImageInfo
	Find(code string) ascii.ImageAscii
	InsertCommand(ascii.ImageAscii)
}