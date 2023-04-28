package utils

import (
	"io"
	"io/ioutil"
	"net/http"
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

func CleanScapeChars(str string) (asciiClean string) {
	asciiClean = strings.Replace(str, "%", "%%", -1)
	return
}