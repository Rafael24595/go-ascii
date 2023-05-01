package sources

import (
	"os"
	"strconv"
	"strings"
)

func BuildBack() string {
	var html strings.Builder
	uri := "/api/view/ascii"
	html.WriteString("<p>")
	html.WriteString("<<< Menu: ")
	html.WriteString("<a href=\"" + uri + "\">" + uri + "</a>")
	html.WriteString("</p>")
	return html.String()
}

func BuildLine() string {
	var html strings.Builder
	html.WriteString("<p style=\"width: 100%; border-bottom: 2px solid;\"></p>")
	return html.String()
}

func BuildReloadScript(ms int) string {
	scriptByte, err := os.ReadFile("src/view/sources/AsciiReloadScript.html")
	if(err != nil){
		panic(err)
	}

	script := string(scriptByte)
	script = strings.Replace(script, "$TIMEOUT", strconv.Itoa(ms), -1)

	return script
}