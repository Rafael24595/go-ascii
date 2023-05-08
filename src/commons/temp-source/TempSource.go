package tempsource

import (
	"os"
	"errors"
	"strings"
	"path/filepath"
	"encoding/base64"
	"golang.org/x/exp/slices"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/utils"
)

const directory = ".temp"
var temps = []string{}

func Base64ToSource(encode string, code string) (path string) {
	createTempDirIfNotExists()

	dec, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		panic(err)
	}

	file := createTemp(code)
	defer file.Close()
	
	if _, err := file.Write(dec); err != nil {
		panic(err)
	}
	if err := file.Sync(); err != nil {
		panic(err)
	}

	path = filepath.ToSlash(file.Name())
	temps = append(temps, path)
	return
}

func createTemp(code string) *os.File {
	name := buildFileName(code)

	file, err := os.CreateTemp(".temp", name)
	if err != nil {
		panic(err)
	}
	path := filepath.ToSlash(file.Name())
	fileCode := filepath.Base(path)
	if isCodePersisted(fileCode) {
		err := os.Remove(path)
		if err != nil {
			panic(err)
		}
		return createTemp(code)
	}
	return file
}

func buildFileName(code string) (name string) {
	code = strings.TrimSpace(code)
	name = "Source-"
	if(code != ""){
		name += code + "-"
	}
	return
}

func isCodePersisted(code string) bool {
	depencencyContainer := dependency_container.GetInstance()
	queryRepository := depencencyContainer.GetQueryRepository()
	image := queryRepository.Find(code)
	return image.GetStatus() != ""
}

func createTempDirIfNotExists() bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)
		if err != nil {
			panic(err)
		}
		return false
	}
	return true
}

func CleanSessionSources() {
	createTempDirIfNotExists()
	for _, temp := range temps {
		err := os.Remove(temp)
		if err != nil {
			panic(err)
		}
	}
}

func CleanSource(code string) {
	if createTempDirIfNotExists() {
		path, err := findSourcePath(code)
		if err != nil {
			panic(err)
		}
		err = os.Remove(path)
		if err != nil {
			panic(err)
		}
		removeSourcePath(code)
	}
}

func findSourcePath(code string) (string, error) {
	idx := slices.IndexFunc(temps, func(path string) bool { return filepath.ToSlash(path) == filepath.ToSlash(directory + "/" + code) })
	if idx != -1 {
		return temps[idx], nil
	} else {
		return "", errors.New("Temp source does not exists.")
	}
}

func removeSourcePath(code string) {
	idx := slices.IndexFunc(temps, func(path string) bool { return filepath.ToSlash(path) == filepath.ToSlash(directory + "/" + code) })
	if idx != -1 {
		temps = utils.RemoveIndex(temps, idx)
	}
}