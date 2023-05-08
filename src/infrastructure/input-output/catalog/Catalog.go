package view_catalog

import "os"

type SourceDirectory string

const (
	directory          string          = "src/infrastructure/input-output/catalog/"
	AsciiDeleteScript  SourceDirectory = "AsciiDeleteScript.html"
	AsciiGifDelay      SourceDirectory = "AsciiGifDelay.html"
	AsciiGifScript     SourceDirectory = "AsciiGifScript.html"
	AsciiLogForm     SourceDirectory = "AsciiLogForm.html"
	AsciiPostForm      SourceDirectory = "AsciiPostForm.html"
	AsciiReloadScript  SourceDirectory = "AsciiReloadScript.html"
	AsciiRestoreScript SourceDirectory = "AsciiRestoreScript.html"
	AsciiTableStyle SourceDirectory = "AsciiTableStyle.html"
)

func GetSource(code SourceDirectory) string {
	path := directory + string(code)
	scriptByte, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(scriptByte)
}