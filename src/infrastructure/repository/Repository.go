package repository

import "go-ascii/src/domain/ascii"

type Repository interface {
	FindAllAscii() string
	FindAscii(code string) ascii.ImageAscii
	InsertAscii(image ascii.ImageAscii) string
}