package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func ParseFloat64(number string) (float64, error) {
	if s, err := strconv.ParseFloat(number, 64); err == nil {
		return s, nil
	} else {
		return 0, err
	}
}

func ParseBoolean(boolean string) (bool, error) {
	if s, err := strconv.ParseBool(boolean); err == nil {
		return s, nil
	} else {
		return false, err
	}
}