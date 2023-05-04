package ascii_view

import (
	"strings"
	"strconv"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/io/sources"
)

type AsciiViewBuilder struct {
	image dto.AsciiResponse
	args map[string]string
}

func NewAsciiViewBuilder(image dto.AsciiResponse, args map[string]string) AsciiViewBuilder {
	return AsciiViewBuilder{image: image, args: args}
}

func (this AsciiViewBuilder) Build() (body string) {
	var html strings.Builder

	html.WriteString(sources.BuildBack())

	html.WriteString(sources.BuildLine())

	html.WriteString(this.buildCode())
	html.WriteString(this.buildType())
	html.WriteString(this.buildScale())
	html.WriteString(this.buildStatus())
	html.WriteString(this.buildMessage())

	if len(this.image.Frames) == 1 {
		static := newAsciiStaticViewBuilder(this.image, this.args)
		html.WriteString(static.Build())
	} else if len(this.image.Frames) > 1 {
		animation := newAsciiAnimationViewBuilder(this.image, this.args)
		html.WriteString(animation.Build())
	}

	if this.image.Status == request_state.PROCESS || this.image.Status == request_state.PENDING {
		html.WriteString(sources.BuildReloadScript(3000))
		html.WriteString("<p><b>*** This page will be reloaded in a few seconds, please wait. ***</b></p>")
	} 

	return html.String()
}

func (this AsciiViewBuilder) buildCode() string {
	var body strings.Builder
	body.WriteString("<p>")
	body.WriteString("Code: " + this.image.Name)
	body.WriteString("</p>")
	return body.String()
}

func (this AsciiViewBuilder) buildType() string {
	var body strings.Builder
	body.WriteString("<p>")
	body.WriteString("Type: " + this.image.Extension)
	body.WriteString("</p>")
	return body.String()
}

func (this AsciiViewBuilder) buildScale() string {
	var body strings.Builder
	body.WriteString("<p>")
	body.WriteString("Height: " + strconv.Itoa(this.image.Height))
	body.WriteString(" - ")
	body.WriteString("Width: " + strconv.Itoa(this.image.Width))
	body.WriteString("</p>")
	return body.String()
}

func (this AsciiViewBuilder) buildStatus() string {
	var body strings.Builder
	body.WriteString("<p>")
	body.WriteString("Status: " + this.image.Status)
	body.WriteString("</p>")
	return body.String()
}

func (this AsciiViewBuilder) buildMessage() string {
	var body strings.Builder
	if(this.image.Message != ""){
		body.WriteString("<p>")
		body.WriteString("Message: " + this.image.Message)
		body.WriteString("</p>")
	}
	return body.String()
}