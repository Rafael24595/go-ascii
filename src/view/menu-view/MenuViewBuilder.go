package menu_view

import (
	"strings"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/view/sources"
)

type MenuViewBuilder struct {
	sources []ascii.ImageInfo
}

func NewMenuViewBuilder(sources []ascii.ImageInfo) MenuViewBuilder {
	return MenuViewBuilder{sources: sources}
}

func (this MenuViewBuilder) Build() (body string) {
	var html strings.Builder

	html.WriteString(this.buildForm())
	html.WriteString(sources.BuildLine())

	if(len(this.sources) == 0){
		html.WriteString("<p>There are no data sources to show.</p>")
	} else {
		html.WriteString("<ul>")

		for _, source := range this.sources {
			uri := "/api/view/ascii/" + source.Code
			html.WriteString("<li>")
			html.WriteString("<b>" + source.Code + " - [" + source.Type + "]</b> >>> ")
			html.WriteString("<a href=\"" + uri + "\">" + uri + "</a>")
			html.WriteString("</li>")
		}

		html.WriteString("</ul>")
	}
	
	return html.String()
}

func (this MenuViewBuilder) buildForm() (body string) {
	var html strings.Builder

	uri := "/api/view/form"
	html.WriteString(">>> Upload file: ")
	html.WriteString("<a href=\"" + uri + "\">" + uri + "</a>")
	
	return html.String()
}