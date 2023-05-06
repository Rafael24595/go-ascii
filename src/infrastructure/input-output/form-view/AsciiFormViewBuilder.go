package form_view

import (
	"os"
	"strings"
	"go-ascii/src/commons/constants/gray-scale"
	"go-ascii/src/infrastructure/input-output/sources"
	"go-ascii/src/infrastructure/input-output/sources/dictionary"
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
	path := dictionary.GetSource(dictionary.AsciiPostForm)
	scriptByte, err := os.ReadFile(path)
	if(err != nil){
		panic(err)
	}

	html := string(scriptByte)
	html = strings.Replace(html, "$GRAYSCALES", this.builGrayScaleOptions(), -1)

	return html
}

func (this AsciiFormViewBuilder) builGrayScaleOptions() string {
	var html strings.Builder
	for _, key := range gray_scale.GetGrayScaleList() {
		selected := ""
		if key == string(gray_scale.DEFAULT) {
			selected = "selected"
		}
		html.WriteString("<option value=\""+ key +"\" "+ selected +">"+ key +"</option>")
	}
	return html.String()
}