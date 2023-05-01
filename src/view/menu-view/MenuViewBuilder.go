package menu_view

import (
	"strings"
)

type MenuViewBuilder struct {
	host string
	sources []string
}

func NewMenuViewBuilder(sources []string, host string) MenuViewBuilder {
	return MenuViewBuilder{sources: sources, host: host}
}

func (this MenuViewBuilder) Build() (body string) {
	var html strings.Builder

	html.WriteString(this.buildForm())

	if(len(this.sources) == 0){
		html.WriteString("<p>There are no data sources to show.</p>")
	} else {
		html.WriteString("<ul>")

		for _, source := range this.sources {
			uri := "/api/view/ascii/" + source
			html.WriteString("<li>")
			html.WriteString("<b>" + source + "</b>: ")
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