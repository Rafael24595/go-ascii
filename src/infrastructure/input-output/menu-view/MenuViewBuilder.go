package menu_view

import (
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/input-output"
	"go-ascii/src/infrastructure/input-output/catalog"
	"strings"
)

type MenuViewBuilder struct {
	sources []dto.InfoResponse
}

func NewMenuViewBuilder(sources []dto.InfoResponse) MenuViewBuilder {
	return MenuViewBuilder{sources: sources}
}

func (this MenuViewBuilder) Build() (body string) {
	var html strings.Builder

	html.WriteString(this.buildForm())
	html.WriteString(input_output.BuildLine())

	if(len(this.sources) == 0){
		html.WriteString("<p>There are no data sources to show.</p>")
	} else {
		html.WriteString("<ul>")

		for _, source := range this.sources {
			uriSource := "/api/ascii/" + source.Code
			uriView := "/api/view/ascii"
			uriViewSource := uriView + "/" + source.Code
			style := ""
			if source.Status == request_state.DELETED {
				style += "text-decoration: line-through;"
			}
			html.WriteString("<li style=\"" + style + "\">")
			html.WriteString("<b>" + source.Code + " - [" + source.Extension + "]</b> >>> ")
			html.WriteString("<ul>")
			html.WriteString("<li>")
			html.WriteString("<span>View: </span>")
			html.WriteString("<a href=\"" + uriViewSource + "\">" + uriViewSource + "</a>")
			html.WriteString("</li>")
			if source.Status != request_state.DELETED {
				html.WriteString("<li>")
				html.WriteString("<span>Delete: </span>")
				html.WriteString("<a onclick=\"deleteAscii(event)\" code=\"" +  source.Code+ "\" view=\"" + uriView + "\" href=\"" + uriSource + "\">" + uriSource + "</a>")
				html.WriteString("</li>")
			}
			html.WriteString("</ul>")
			html.WriteString("</li>")
			html.WriteString("<br>")
		}

		html.WriteString("</ul>")
	}

	html.WriteString(this.buildDeleteScript())
	
	return html.String()
}

func (this MenuViewBuilder) buildForm() (body string) {
	configuration := configuration.GetInstance()

	var html strings.Builder

	uriForm := "/api/view/form"
	html.WriteString("<span>>>> Upload file: ")
	html.WriteString("<a href=\"" + uriForm + "\">" + uriForm + "</a></span>")

	if configuration.IsDebugSession() {
		logUri := "/api/view/log"
		html.WriteString("<span style=\"float:right;\">>>> :Session log: ")
		html.WriteString("<a href=\"" + logUri + "\">" + logUri + "</a></span>")
	}
	
	return html.String()
}

func (this MenuViewBuilder) buildDeleteScript() string {
	return view_catalog.GetSource(view_catalog.AsciiDeleteScript)
}