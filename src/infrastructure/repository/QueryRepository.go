package repository

import "go-ascii/src/domain/ascii"

type QueryRepository interface {
	FindAllAscii() []string
	FindAscii(code string) ascii.ImageAscii
	InsertCommand(ascii.ImageAscii)
}