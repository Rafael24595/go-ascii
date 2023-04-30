package ascii_view

import (
	"strings"
	"go-ascii/src/domain/ascii"
)

type asciiStaticViewBuilder struct {
	image ascii.ImageAscii
}

func newAsciiStaticViewBuilder(image ascii.ImageAscii) asciiStaticViewBuilder {
	return asciiStaticViewBuilder{image: image}
}

func (this asciiStaticViewBuilder) Build() string {
	var html strings.Builder

	html.WriteString(this.buildStaticBody())
	
	return html.String()
}

func (this asciiStaticViewBuilder) buildStaticBody() string {
	var body strings.Builder

	body.WriteString("<pre id=\"" + this.image.Name + "\" type=\"" + this.image.Type + "\">")
	body.WriteString(this.image.Frames[0])
	body.WriteString("</pre>")

	return body.String()
}