package utils

import (
	"os"
	"io"
	"bytes"
	"strings"
	"net/http"
	"io/ioutil"
)

func ReaderExtension(reader io.Reader) string {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return http.DetectContentType(data)
}

func FileExtensionByFile(file *os.File) string {
	var buffer bytes.Buffer
    reader := io.TeeReader(file, &buffer)
	extension := ReaderExtension(reader)
	return strings.Split(extension, ";")[0]
}

func FileExtensionByPath(path string) (extesion string) {
	temp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	extesion = FileExtensionByFile(temp)
	temp.Close()
	return
}

func CleanScapeChars(str string) string {
	return strings.Replace(str, "%", "%%", -1)
}

func RemoveIndex[T interface{}](s []T, index int) []T{
	return append(s[:index], s[index+1:]...)
}