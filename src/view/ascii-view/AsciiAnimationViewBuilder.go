package ascii_view

import (
	"os"
	"strconv"
	"strings"
	"go-ascii/src/domain/ascii"
)

type asciiAnimationViewBuilder struct {
	image ascii.ImageAscii
}

func newAsciiAnimationViewBuilder(image ascii.ImageAscii) asciiAnimationViewBuilder {
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

	body.WriteString("<pre id=\"" + this.image.Name + "\" type=\"" + this.image.Type + "\">")
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