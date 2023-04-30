package image

import (
	"io"
	"os"
	"bytes"
	"image/gif"
	"image/png"
	"encoding/base64"
	"go-ascii/src/commons/utils"
)

func Encoder(file *os.File) (encode string) {
	var buffer bytes.Buffer
    reader := io.TeeReader(file, &buffer)
	extension := utils.ReaderExtension(reader)
	encode = encodeByExtension(buffer, extension)
	return
}

func encodeByExtension(buffer bytes.Buffer, extension string) (encode string) {
	var buf bytes.Buffer
	switch extension {
		case "image/jpeg", "image/jpg":
			buf = encodeJpg(buffer)
		case "image/gif":
			buf = encodeGif(buffer)
		case "image/png":
			buf = encodePng(buffer)
		default:
			panic("unknown file type uploaded")
	}
	data := buf.Bytes()
	encode = base64.StdEncoding.EncodeToString(data)
	return
}

func encodeJpg(buffer bytes.Buffer) (buf bytes.Buffer) {
	decode, err := gif.DecodeAll(&buffer)
	if err != nil {
		panic(err)
	}
	if err := gif.EncodeAll(&buf, decode); err != nil {
		panic(err)
	}
	return
}

func encodeGif(buffer bytes.Buffer) (buf bytes.Buffer) {
	decode, err := gif.DecodeAll(&buffer)
	if err != nil {
		panic(err)
	}
	if err := gif.EncodeAll(&buf, decode); err != nil {
		panic(err)
	}
	return
}

func encodePng(buffer bytes.Buffer) (buf bytes.Buffer) {
	decode, err := png.Decode(&buffer)
	if err != nil {
		panic(err)
	}
	if err := png.Encode(&buf, decode); err != nil {
		panic(err)
	}
	return
}