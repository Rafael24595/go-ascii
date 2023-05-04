package dictionary

type SourceDirectory string

const (
	directory SourceDirectory = "src/infrastructure/io/sources/"
	AsciiGifDelay SourceDirectory = "AsciiGifDelay.html"
	AsciiGifScript SourceDirectory = "AsciiGifScript.html"
	AsciiPostForm SourceDirectory = "AsciiPostForm.html"
	AsciiReloadScript SourceDirectory = "AsciiReloadScript.html"
)

func GetSource(code SourceDirectory) string {
	return string(directory) + string(code)
}