package tempsource

import (
	"os"
	"errors"
	"strings"
	"path/filepath"
	"encoding/base64"
	"golang.org/x/exp/slices"
	"go-ascii/src/commons/constants/log-categories"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log"
	"go-ascii/src/commons/utils"
)

const directory = ".temp"
var temps = []string{}

func Base64ToSource(encode string, code string) (path string) {
	_ = createTempDirIfNotExists()

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

	file, err := os.CreateTemp(directory, name)
	if err != nil {
		panic(err)
	}
	path := filepath.ToSlash(file.Name())
	fileCode := filepath.Base(path)
	log.Log(log_categories.INFO, "Temporal file \"" + fileCode + "\" created.")
	if isCodePersisted(fileCode) {
		log.Log(log_categories.WARNING, "Registry \"" + fileCode + "\" already exists. Recreating it with another code...")
		removeSourceFile(path)
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
	image, ok := queryRepository.Find(code)
	return ok && image.GetStatus() != ""
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
	if createTempDirIfNotExists() {
		for _, temp := range temps {
			removeSourceFile(temp)
		}
	}
}

func CleanSource(code string) {
	if createTempDirIfNotExists() {
		path, err := findSourcePath(code)
		if err != nil {
			panic(err)
		}
		removeSourceFile(path)
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

func removeSourceFile(temp string) {
	fileCode := filepath.Base(temp)
	err := os.Remove(temp)
		if err != nil {
			log.Log(log_categories.ERROR, "Cannot remove file \"" + fileCode + "\".")
			panic(err)
	}
	log.Log(log_categories.INFO, "Temporal file \"" + fileCode + "\" deleted.")
}