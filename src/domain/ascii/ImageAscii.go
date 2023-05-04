package ascii

import (
	"strings"
)

type ImageAscii struct {
	name      string
	extension string
	status    string
	frames    []string
}

func NewImageAscii(name string, extension string, status string, frames []string) ImageAscii {
	return ImageAscii{name: name, extension: extension, status: status, frames: frames}
}

func (this *ImageAscii) GetName() string {
	return this.name
}

func (this *ImageAscii) GetExtension() string {
	return this.extension
}

func (this *ImageAscii) GetStatus() string {
	return this.status
}

func (this *ImageAscii) GetFrames() []string {
	return this.frames
}

func (this *ImageAscii) SetName(name string) {
	this.name = name
}

func (this *ImageAscii) SetExtension(extension string) {
	this.extension = extension
}

func (this *ImageAscii) SetStatus(status string) {
	this.status = status
}

func (this *ImageAscii) AppendFrame(frame string) {
	this.frames = append(this.frames, frame)
}

func (this *ImageAscii) GetDimensions() (int, int) {
	rowsCount := 0
	rowsLength := 0
	for _, frame := range this.frames {
		rows := strings.Split(frame, "\n")
		for _, row := range rows {
			if len(row) > 0 {
				rowsLength += len(row)
				rowsCount++
			}
		}
	}
	if(rowsCount != 0 && rowsLength != 0){
		height := rowsCount / len(this.frames)
		width := (rowsLength / len(this.frames)) / height
		return height , width
	} else {
		return 0, 0
	}
}