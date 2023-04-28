package image

import (
	"bytes"
	"go-ascii/src/commons/utils"
	"image"
	"image/png"
	"io"
	"os"
)

func Decoder(file *os.File) (decode image.Image) {
	var buffer bytes.Buffer
    reader := io.TeeReader(file, &buffer)
	extension := utils.ReaderExtension(reader)
	decode = decodeByExtension(buffer, extension)
	return
}


func decodeByExtension(buffer bytes.Buffer, extension string) (decode image.Image) {
	var err error
	switch extension {
		case "image/jpeg", "image/jpg":
			decode, err = png.Decode(&buffer)
		case "image/gif":
			decode, err = png.Decode(&buffer)
		case "image/png":
			decode, err = png.Decode(&buffer)
		default:
			panic("unknown file type uploaded")
	}

	if err != nil {
		panic(err)
	}

	return
}