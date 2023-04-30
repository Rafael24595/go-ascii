package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ReaderExtension(reader io.Reader) string {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return http.DetectContentType(data)
}

func FileExtension(file *os.File) string {
	var buffer bytes.Buffer
    reader := io.TeeReader(file, &buffer)
	extension := ReaderExtension(reader)
	return strings.Split(extension, ";")[0]
}

func CleanScapeChars(str string) string {
	return strings.Replace(str, "%", "%%", -1)
}