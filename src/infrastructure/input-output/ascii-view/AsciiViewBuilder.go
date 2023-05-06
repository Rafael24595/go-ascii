package ascii_view

import (
	"os"
	"strconv"
	"strings"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/input-output/sources"
	"go-ascii/src/infrastructure/input-output/sources/dictionary"
)

type AsciiViewBuilder struct {
	image dto.InfoAsciiResponse
	args map[string]string
}

func NewAsciiViewBuilder(image dto.InfoAsciiResponse, args map[string]string) AsciiViewBuilder {
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

	if this.image.Status == request_state.DELETED {
		uriSource := "/api/ascii/" + this.image.Name
		uriView := "/api/view/ascii"
		html.WriteString("<p><b>*** This source will be deleted, click <a onclick=\"restoreAscii(event)\" code=\"" +  this.image.Name + "\" view=\"" + uriView + "\" href=\"" + uriSource + "\">HERE</a> to restore it. ***</b></p>")
		html.WriteString(this.buildRestoreScript())
	} 

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

func (this AsciiViewBuilder) buildRestoreScript() string {
	path := dictionary.GetSource(dictionary.AsciiRestoreScript)
	scriptByte, err := os.ReadFile(path)
	if(err != nil){
		panic(err)
	}

	return string(scriptByte)
}