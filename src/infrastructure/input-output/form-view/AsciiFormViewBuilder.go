package form_view

import (
	"strings"
	"go-ascii/src/commons/constants/gray-scale"
	"go-ascii/src/infrastructure/input-output/sources"
	"go-ascii/src/infrastructure/input-output/sources/catalog"
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
	html := view_catalog.GetSource(view_catalog.AsciiPostForm)
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