package repository

import "go-ascii/src/domain/ascii"

type CommandRepository interface {
	InsertAscii(image ascii.ImageAscii) string
	InsertQuery(image ascii.ImageAscii)
}