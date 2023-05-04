package repository

import (
	"go-ascii/src/commons"
	"go-ascii/src/domain/ascii"
)

type CommandRepository interface {
	commons.Dependency
	InsertAscii(image ascii.ImageAscii) string
	InsertQuery(image ascii.ImageAscii)
}