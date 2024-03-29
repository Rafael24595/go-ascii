package ascii_view

import (
	"strings"
	"go-ascii/src/commons/dto"
)

type asciiStaticViewBuilder struct {
	image dto.InfoAsciiResponse
	args map[string]string
}

func newAsciiStaticViewBuilder(image dto.InfoAsciiResponse, args map[string]string) asciiStaticViewBuilder {
	return asciiStaticViewBuilder{image: image, args: args}
}

func (this asciiStaticViewBuilder) Build() string {
	var html strings.Builder

	html.WriteString(this.buildStaticBody())
	
	return html.String()
}

func (this asciiStaticViewBuilder) buildStaticBody() string {
	var body strings.Builder

	body.WriteString("<pre id=\"" + this.image.Name + "\" type=\"" + this.image.Extension + "\">")
	body.WriteString(this.image.Frames[0])
	body.WriteString("</pre>")

	return body.String()
}