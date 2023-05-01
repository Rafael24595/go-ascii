package ascii

type ImageInfo struct {
	code string
	extension string
}

func NewImageInfo(code string, extension string) ImageInfo {
	return ImageInfo{code: code, extension: extension}
}

func (this ImageInfo) GetCode() string {
	return this.code
}

func (this ImageInfo) GetExtension() string {
	return this.extension
}