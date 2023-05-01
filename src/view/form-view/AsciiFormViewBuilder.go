package form_view

import (
	"os"
	"strings"
	"go-ascii/src/view/sources"
)

type AsciiFormViewBuilder struct {
}

func NewAsciiAsciiFormViewBuilder() AsciiFormViewBuilder {
	return AsciiFormViewBuilder{}
}

func (this AsciiFormViewBuilder) Build() string {
	var html strings.Builder
	html.WriteString(sources.BuildBack())
	html.WriteString(sources.BuildLine())
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