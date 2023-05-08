package ascii_view

import (
	"strconv"
	"strings"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/input-output/catalog"
)

type asciiAnimationViewBuilder struct {
	image dto.InfoAsciiResponse
	args map[string]string
}

func newAsciiAnimationViewBuilder(image dto.InfoAsciiResponse, args map[string]string) asciiAnimationViewBuilder {
	return asciiAnimationViewBuilder{image: image, args: args}
}

func (this asciiAnimationViewBuilder) Build() string {
	var html strings.Builder

	html.WriteString(this.buildDelayForm())
	html.WriteString(this.buildAnimationBody())
	html.WriteString(this.buildAnimationScript())

	return html.String()
}

func (this asciiAnimationViewBuilder) buildAnimationBody() string {
	var body strings.Builder

	body.WriteString("<pre id=\"" + this.image.Name + "\" type=\"" + this.image.Extension + "\">")
	for i, frame := range this.image.Frames {
		body.WriteString("<pre id=\"" + this.image.Name + "-" + strconv.Itoa(i) + "\" style=\"display: none;\">")
		body.WriteString(frame)
		body.WriteString("</pre>")
	}
	body.WriteString("</pre>")

	return body.String()
}

func (this asciiAnimationViewBuilder) buildDelayForm() string {
	return view_catalog.GetSource(view_catalog.AsciiGifDelay)
}

func (this asciiAnimationViewBuilder) buildAnimationScript() string {
	delay := this.args["delay"]
	if delay == "" {
		delay = "250"
	}
	
	script := view_catalog.GetSource(view_catalog.AsciiGifScript)
	script = strings.Replace(script, "$GIFS", "\"" + this.image.Name + "\"", -1)
	script = strings.Replace(script, "$DELAY", "\"" + delay + "\"", -1)

	return script
}