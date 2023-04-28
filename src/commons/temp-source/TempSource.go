package tempsource

import (
	"encoding/base64"
	"os"
)

var temps = []string{}

func Base64ToSource(encode string) (path string) {
	dec, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		panic(err)
	}

	file, err := os.CreateTemp(".temp", "Source-")
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

func CleanSessionSources() {
	for _, temp := range temps {
		err := os.Remove(temp)
		if err != nil {
			panic(err)
		}
	}
}