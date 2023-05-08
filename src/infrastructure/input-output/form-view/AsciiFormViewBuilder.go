package form_view

import (
	"go-ascii/src/commons/constants/gray-scale"
	"go-ascii/src/infrastructure/input-output"
	"go-ascii/src/infrastructure/input-output/catalog"
	"strings"
)

type AsciiFormViewBuilder struct {
}

func NewAsciiAsciiFormViewBuilder() AsciiFormViewBuilder {
	return AsciiFormViewBuilder{}
}

func (this AsciiFormViewBuilder) Build() string {
	var html strings.Builder
	html.WriteString(input_output.BuildBack())
	html.WriteString(input_output.BuildLine())
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