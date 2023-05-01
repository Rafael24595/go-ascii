package form_view

import (
	ascii_view "go-ascii/src/view/ascii-view"
	"os"
	"strings"
)

type AsciiFormViewBuilder struct {
}

func NewAsciiAsciiFormViewBuilder() AsciiFormViewBuilder {
	return AsciiFormViewBuilder{}
}

func (this AsciiFormViewBuilder) Build() string {
	var html strings.Builder
	html.WriteString(ascii_view.BuildBack())
	html.WriteString(this.buildFormScript())
	return html.String()
}

func (this AsciiFormViewBuilder) buildFormScript() string {
	scriptByte, err := os.ReadFile("src/view/sources/AsciiPostForm.html")
	if(err != nil){
		panic(err)
	}

	return string(scriptByte)
}