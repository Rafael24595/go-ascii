package encoder_decoder

import (
	"bytes"
	"errors"
	"go-ascii/src/commons/utils"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func DecodeByPath(path string) (decode []image.Image, err error) {
	temp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	decode, err = DecodeByFile(temp)
	temp.Close()
	return
}

func DecodeByFile(file *os.File) (decode []image.Image, err error) {
	var buffer bytes.Buffer
    reader := io.TeeReader(file, &buffer)
	extension := utils.ReaderExtension(reader)
	decode, err =decodeByExtension(buffer, extension)
	return
}

func decodeByExtension(buffer bytes.Buffer, extension string) (decode []image.Image, err error) {
	switch extension {
		case "image/jpeg", "image/jpg":
			decode = decodeJpg(buffer)
		case "image/gif":
			decode = decodeGif(buffer)
		case "image/png":
			decode = decodePng(buffer)
		default:
			err = errors.New("Unknown file type uploaded")
	}
	return
}

func decodeJpg(buffer bytes.Buffer) (decode []image.Image) {
	decodePng, err := jpeg.Decode(&buffer)
	if err != nil {
		panic(err)
	}
	decode = append(decode, decodePng)
	return
}

func decodeGif(buffer bytes.Buffer) (decode []image.Image) {
	decodeGif, err := gif.DecodeAll(&buffer)
	if err != nil {
		panic(err)
	}
	for _, decodeFrame := range decodeGif.Image {
		decodeFrameAsImg := decodeFrame.SubImage(decodeFrame.Bounds())
		decode = append(decode, decodeFrameAsImg)
	}
	return
}

func decodePng(buffer bytes.Buffer) (decode []image.Image) {
	decodePng, err := png.Decode(&buffer)
	if err != nil {
		panic(err)
	}
	decode = append(decode, decodePng)
	return
}