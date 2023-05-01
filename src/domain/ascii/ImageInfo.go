package ascii

type ImageInfo struct {
	Code string
	Type string
}

func NewImageInfo(code string, typ string) ImageInfo {
	return ImageInfo{Code: code, Type: typ}
}