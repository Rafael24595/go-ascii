package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ReaderExtension(reader io.Reader) (extension string) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	extension = http.DetectContentType(data)
	return
}

func FileExtension(file *os.File) (extension string) {
	var buffer bytes.Buffer
    reader := io.TeeReader(file, &buffer)
	extension = ReaderExtension(reader)
	extension = strings.Split(extension, ";")[0]
	return
}

func CleanScapeChars(str string) (asciiClean string) {
	asciiClean = strings.Replace(str, "%", "%%", -1)
	return
}