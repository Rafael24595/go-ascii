package repository

import (
	"go-ascii/src/commons"
	"go-ascii/src/domain/ascii"
)

type CommandRepository interface {
	commons.Dependency
	Insert(image ascii.ImageAscii) string
	Modify(image ascii.ImageAscii) string
	Delete(image ascii.ImageAscii) string
	//@Deprecated
	ToQuery(image ascii.ImageAscii)
}