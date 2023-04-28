package image

import (
	"bytes"
	"encoding/base64"
	"go-ascii/src/commons/utils"
	"image/png"
	"io"
	"os"
)

func Encoder(file *os.File) (encode string) {
	var buffer bytes.Buffer
    reader := io.TeeReader(file, &buffer)
	extension := utils.ReaderExtension(reader)
	encode = encodeByExtension(buffer, extension)
	return
}

func encodeByExtension(buffer bytes.Buffer, extension string) (encode string) {
	var data []byte
	var buf bytes.Buffer

	decode := decodeByExtension(buffer, extension)
	switch extension {
		case "image/jpeg", "image/jpg":
			if err := png.Encode(&buf, decode); err != nil {
				panic(err)
			}
		case "image/gif":
			if err := png.Encode(&buf, decode); err != nil {
				panic(err)
			}
		case "image/png":
			if err := png.Encode(&buf, decode); err != nil {
				panic(err)
			}
		default:
			panic("unknown file type uploaded")
	}

	data = buf.Bytes()
	encode = base64.StdEncoding.EncodeToString(data)

	return
}