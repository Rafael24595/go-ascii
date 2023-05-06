package ascii

type ImageInfo struct {
	code string
	status string
	extension string
}

func NewImageInfo(code string, status string, extension string) ImageInfo {
	return ImageInfo{code: code, status: status, extension: extension}
}

func (this ImageInfo) GetCode() string {
	return this.code
}

func (this ImageInfo) GetStatus() string {
	return this.status
}

func (this ImageInfo) GetExtension() string {
	return this.extension
}