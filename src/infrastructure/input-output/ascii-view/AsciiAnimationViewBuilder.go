package ascii_view

import (
	"os"
	"strconv"
	"strings"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/input-output/sources/dictionary"
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
	scriptByte, err := os.ReadFile(dictionary.GetSource(dictionary.AsciiGifDelay))
	if(err != nil){
		panic(err)
	}

	return string(scriptByte)
}

func (this asciiAnimationViewBuilder) buildAnimationScript() string {
	scriptByte, err := os.ReadFile(dictionary.GetSource(dictionary.AsciiGifScript))
	if(err != nil){
		panic(err)
	}

	delay := this.args["delay"]
	if delay == "" {
		delay = "250"
	}
	
	script := string(scriptByte)
	script = strings.Replace(script, "$GIFS", "\"" + this.image.Name + "\"", -1)
	script = strings.Replace(script, "$DELAY", "\"" + delay + "\"", -1)

	return script
}