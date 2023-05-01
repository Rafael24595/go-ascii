package ascii_view

import (
	"os"
	"strconv"
	"strings"
	"go-ascii/src/infrastructure/dto"
)

type asciiAnimationViewBuilder struct {
	image dto.AsciiResponse
}

func newAsciiAnimationViewBuilder(image dto.AsciiResponse) asciiAnimationViewBuilder {
	return asciiAnimationViewBuilder{image: image}
}

func (this asciiAnimationViewBuilder) Build() string {
	var html strings.Builder

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

func (this asciiAnimationViewBuilder) buildAnimationScript() string {
	scriptByte, err := os.ReadFile("src/view/sources/AsciiGifScript.html")
	if(err != nil){
		panic(err)
	}

	script := string(scriptByte)
	script = strings.Replace(script, "$GIFS", "\"" + this.image.Name + "\"", -1)

	return script
}