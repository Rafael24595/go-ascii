package log_view

import (
	"time"
	"strings"
	"strconv"
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/constants/log-categories"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/utils"
	"go-ascii/src/infrastructure/input-output"
	"go-ascii/src/infrastructure/input-output/catalog"
)

type LogViewBuilder struct {
	sources []dto.InfoLogResponse
}

func NewLogViewBuilder(sources []dto.InfoLogResponse) LogViewBuilder {
	return LogViewBuilder{sources: sources}
}

func (this LogViewBuilder) Build(dto dto.LogParamsRequest) (body string) {
	var html strings.Builder

	html.WriteString(input_output.BuildBack())
	html.WriteString(input_output.BuildLine())

	html.WriteString(this.buildSessionData())

	html.WriteString(this.buildFormScript(dto))

	if(len(this.sources) == 0){
		html.WriteString("<p>There are no log sources to show.</p>")
	} else {
		layout := "2006-01-02_15:04:05"

		html.WriteString(input_output.BuildTableStyles())

		html.WriteString("<table>")
		html.WriteString("<tr><th>Index</th><th>Category</th><th>Family</th><th>Timestamp</th><th>Message</th></tr>")

		for i, source := range this.sources {
			timestamp := time.Unix(0, int64(source.Timestamp) * int64(time.Millisecond))
			cls := this.getStyleClass(source.Category)
			html.WriteString("<tr>")
			html.WriteString("<td class=\"text-centre\"><b>" + strconv.Itoa(i+1) + ".</b></td>")
			html.WriteString("<td class=\"text-centre "+ cls +"\">" + source.Category + "</td>")
			html.WriteString("<td class=\"text-centre\">" + source.Family + "</td>")
			html.WriteString("<td class=\"text-centre\">" + timestamp.Format(layout) + "</td>")
			html.WriteString("<td>" + source.Message + "</td>")
			html.WriteString("</tr>")
		}

		html.WriteString("</table>")
	}
	
	return html.String()
}

func (this LogViewBuilder) getStyleClass(category string) string {
	switch category {
		case string(log_categories.INFO):
			return "success"
		case string(log_categories.WARNING):
			return "warn"
		case string(log_categories.ERROR):
			return "error"
		default:
			return ""
	}
}

func (this LogViewBuilder) buildSessionData() (body string) {
	configuration := configuration.GetInstance()
	layout := "2006-01-02 15:04:05"

	var html strings.Builder

	html.WriteString("<p><b>Session id</b>: " + configuration.GetSessionId()  + "</p>")
	html.WriteString("<p><b>Date</b>: " + configuration.GetTimestamp().Format(layout)  + "</p>")

	return html.String()
}

func (this LogViewBuilder) buildFormScript(dto dto.LogParamsRequest) string {
	html := view_catalog.GetSource(view_catalog.AsciiLogForm)
	html = strings.Replace(html, "$OPT_CATEGORIES", this.builCategoriesOptions(), -1)
	html = strings.Replace(html, "$QUERYPARMS", this.builCacheParams(dto), -1)

	return html
}

func (this LogViewBuilder) builCategoriesOptions() string {
	var html strings.Builder
	for _, key := range log_categories.LogCategoryList() {
		html.WriteString("<option value=\""+ key +"\">"+ key +"</option>")
	}
	return html.String()
}

func (this LogViewBuilder) builCacheParams(dto dto.LogParamsRequest) string {
	layout := "2006-01-02 15:04:05"
	var html strings.Builder
	if dto.Category != ""{
		html.WriteString("\"category\":\"" + dto.Category + "\",")
	}
	if dto.Family != ""{
		html.WriteString("\"family\":\"" + dto.Family + "\",")
	}
	if dto.From != ""{
		ms, _ := utils.ParseInt64(dto.From)
		timestamp := time.Unix(0, int64(ms) * int64(time.Millisecond))
		html.WriteString("\"from\":\"" + timestamp.Format(layout) + "\",")
	}
	if dto.To != ""{
		ms, _ := utils.ParseInt64(dto.To)
		timestamp := time.Unix(0, int64(ms) * int64(time.Millisecond))
		html.WriteString("\"to\":\"" + timestamp.Format(layout) + "\"")
	}
	
	return html.String()
}