package tempsource

import (
	"encoding/base64"
	"os"
	"strings"
)

var temps = []string{}

func Base64ToSource(encode string, code string) (path string) {
	dec, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		panic(err)
	}

	code = strings.TrimSpace(code)
	name := "Source-"
	if(code != ""){
		name += code + "-"
	}

	createTempDirIfNotExists()

	file, err := os.CreateTemp(".temp", name)
	if err != nil {
		panic(err)
	}
	
	if _, err := file.Write(dec); err != nil {
		panic(err)
	}

	if err := file.Sync(); err != nil {
		panic(err)
	}

	defer file.Close()

	path = file.Name()
	temps = append(temps, path)
	return
}

func createTempDirIfNotExists() {
	if _, err := os.Stat(".temp"); os.IsNotExist(err) {
		err := os.Mkdir(".temp", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
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