package ascii

import "time"

type ImageInfo struct {
	code string
	status string
	timestamp time.Time
	extension string
}

func NewImageInfo(code string, status string, timestamp time.Time, extension string) ImageInfo {
	return ImageInfo{code: code, status: status, timestamp: timestamp, extension: extension}
}

func (this ImageInfo) GetCode() string {
	return this.code
}

func (this ImageInfo) GetStatus() string {
	return this.status
}

func (this ImageInfo) GetTimestamp() time.Time {
	return this.timestamp
}

func (this ImageInfo) GetExtension() string {
	return this.extension
}