package ascii_view

import (
	"strings"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/infrastructure/dto"
	"go-ascii/src/view/sources"
)

type AsciiViewBuilder struct {
	image dto.AsciiResponse
}

func NewAsciiViewBuilder(image dto.AsciiResponse) AsciiViewBuilder {
	return AsciiViewBuilder{image: image}
}

func (this AsciiViewBuilder) Build() (body string) {
	var html strings.Builder

	html.WriteString(sources.BuildBack())

	html.WriteString(sources.BuildLine())

	html.WriteString(this.buildCode())
	html.WriteString(this.buildType())
	html.WriteString(this.buildStatus())

	if len(this.image.Frames) == 1 {
		static := newAsciiStaticViewBuilder(this.image)
		html.WriteString(static.Build())
	} else if len(this.image.Frames) > 1 {
		animation := newAsciiAnimationViewBuilder(this.image)
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

func (this AsciiViewBuilder) buildStatus() string {
	var body strings.Builder
	body.WriteString("<p>")
	body.WriteString("Status: " + this.image.Status)
	body.WriteString("</p>")
	return body.String()
}