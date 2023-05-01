package ascii_view

import (
	"go-ascii/src/domain/ascii"
	"strings"
)

type AsciiViewBuilder struct {
	image ascii.ImageAscii
}

func NewAsciiViewBuilder(image ascii.ImageAscii) AsciiViewBuilder {
	return AsciiViewBuilder{image: image}
}

func (this AsciiViewBuilder) Build() (body string) {

	var html strings.Builder

	html.WriteString(BuildBack())

	html.WriteString(this.buildCode())
	html.WriteString(this.buildType())

	if len(this.image.Frames) == 1 {
		static := newAsciiStaticViewBuilder(this.image)
		html.WriteString(static.Build())
	} else if len(this.image.Frames) > 1 {
		animation := newAsciiAnimationViewBuilder(this.image)
		html.WriteString(animation.Build())
	}

	return html.String()
}

func BuildBack() string {
	var html strings.Builder
	uri := "/api/view/ascii"
	html.WriteString("<p>")
	html.WriteString("<<< Menu: ")
	html.WriteString("<a href=\"" + uri + "\">" + uri + "</a>")
	html.WriteString("</p>")
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
	body.WriteString("Type: " + this.image.Type)
	body.WriteString("</p>")
	return body.String()
}