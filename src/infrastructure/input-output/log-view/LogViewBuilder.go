package log_view

import (
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/input-output"
	"strings"
)

type LogViewBuilder struct {
	sources []dto.InfoLogResponse
}

func NewLogViewBuilder(sources []dto.InfoLogResponse) LogViewBuilder {
	return LogViewBuilder{sources: sources}
}

func (this LogViewBuilder) Build() (body string) {
	var html strings.Builder

	html.WriteString(input_output.BuildBack())
	html.WriteString(input_output.BuildLine())

	html.WriteString(this.buildSessionData())

	if(len(this.sources) == 0){
		html.WriteString("<p>There are no log sources to show.</p>")
	} else {
		layout := "2006-01-02_15:04:05"

		html.WriteString(input_output.BuildTableStyles())

		html.WriteString("<table>")
		html.WriteString("<tr><th>Category</th><th>Timestamp</th><th>Message</th></tr>")

		for _, source := range this.sources {
			html.WriteString("<tr>")
			html.WriteString("<td class=\"text-centre\">" + source.Category + "</td>")
			html.WriteString("<td class=\"text-centre\">" + source.Timestamp.Format(layout) + "</td>")
			html.WriteString("<td>" + source.Message + "</td>")
			html.WriteString("</tr>")
		}

		html.WriteString("</table>")
	}
	
	return html.String()
}

func (this LogViewBuilder) buildSessionData() (body string) {
	configuration := configuration.GetInstance()
	layout := "2006-01-02 15:04:05"

	var html strings.Builder

	html.WriteString("<p><b>Session id</b>: " + configuration.GetSessionId()  + "</p>")
	html.WriteString("<p><b>Date</b>: " + configuration.GetTimestamp().Format(layout)  + "</p>")

	return html.String()
}