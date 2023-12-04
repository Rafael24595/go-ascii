package builder

import (
	"go-ascii/src/commons/constants/gray-scale"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/utils"
	"go-ascii/src/commons/utils/encoder-decoder"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder/collection"
	"go-ascii/src/domain/ascii/builder/scale"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strconv"
	"strings"
	"time"
)

//TODO: Parametrize.
const MAX_PERCENTAGE = 0.5

type BuilderMstAscii struct {
	event ascii.QueueEvent
	images collection.ImagesCollection
	imgeScale scale.ImageScale
	grayScale gray_scale.GrayScale
	missingKeys map[string]int
}

func NewBuilderMstAscii(event ascii.QueueEvent) (builder BuilderMstAscii, err error) {
	scaleHeight := event.GetHeight()
	scaleWidth := event.GetWidth()
	grayScale := event.GetGrayScale()

	images, err := encoder_decoder.DecodeByPath(event.GetPath())
	collection := collection.NewImagesCollection(images)
	scale := scale.NewImageScale(collection, scaleHeight, scaleWidth)
	builder = BuilderMstAscii{event: event, images:collection, imgeScale: scale, grayScale: grayScale, missingKeys: map[string]int{}}
	return 
}

func (this BuilderMstAscii) Build() ascii.ImageAscii {
	code := this.event.GetCode()
	extension := utils.FileExtensionByPath(this.event.GetPath())
	imageAscii := ascii.NewImageAscii(code, extension, request_state.SUCCES, time.Now(), []string{})
	for i := range this.images.GetImages() {
		frame := this.buildFrame(i)
		imageAscii.AppendFrame(frame)
	}
	return imageAscii
}

func (this BuilderMstAscii) buildFrame(position int) (frame string) {
	grayscale := this.binarizeFrame(position)

	width := int(math.Floor(float64(grayscale.Rect.Dx()) / 8))
	height := int(math.Floor(float64(grayscale.Rect.Dy()) / 16)) 
	if width <= 1 && height <= 1 {
		panic("Image is too small")
	} 

	for x := 0; x < width; x++ {
		chars := []string{}
		for y := 0; y < height; y++ {
			char := this.chunkCharacter(x, y , grayscale)
			chars = append(chars, char)
		}
		line := strings.Join(chars, "") + "\n";
		frame += string(line)
		print(line)
	}

	//TODO: ¿Finality?
	this.treatMissingKeys()

	return
}

func (this BuilderMstAscii) binarizeFrame(position int) *image.Gray {
	grayscale := this.desaturateImage(position)

	//TODO: ¿"mean" meaning?
	mean := 0
	for x := 0; x < grayscale.Rect.Dx(); x++ {
		for y := 0; y < grayscale.Rect.Dx(); y++ {
			c := grayscale.GrayAt(int(x), int(y)).Y
			mean = mean + int(c)
		}
	}
	mean = mean / (grayscale.Rect.Dx() * grayscale.Rect.Dy())

	for x := 0; x < grayscale.Rect.Dx(); x++ {
		for y := 0; y < grayscale.Rect.Dx(); y++ {
			c := grayscale.GrayAt(int(x), int(y)).Y
			if int(c) > mean {
				grayscale.Set(x, y, color.Gray{0})
			} else {
				grayscale.Set(x, y, color.Gray{1})
			}
		}
	}

	return grayscale
}

func (this BuilderMstAscii) resizeImage(position int) (resized *image.RGBA) {
	height := int(this.images.GetImageHeight())
	width := int(this.images.GetImageWidth())
	resized = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(resized, resized.Bounds(), this.images.GetImage(position), image.Point{0, 0}, draw.Src)
	return
}

func (this BuilderMstAscii) desaturateImage(position int) (grayscale *image.Gray) {
	resized := this.resizeImage(position)
	height := int(this.images.GetImageHeight())
	width := int(this.images.GetImageWidth())

	grayscale = image.NewGray(resized.Bounds())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			grayscale.Set(x, y, c)
		}
	}

	return
}

func (this BuilderMstAscii) chunkCharacter(x int, y int, grayscale *image.Gray) string {
	values := []string{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 16; j++ {
			c := grayscale.GrayAt(int(x*8*i), int(y*16*j)).Y
			values = append(values, strconv.Itoa(int(c)))
		}
	}
	return this.findCharacter(strings.Join(values, ""))
}

func (this BuilderMstAscii) findCharacter(key string) string {
	value := BLANK_CHARACTER

	for char, _ := range CHARACTERS {
		percentage := this.similarity(key, char)
		if percentage > MAX_PERCENTAGE {
			//max_percentage = percentage
			value = CHARACTERS[char]
		}
	}

	//TODO: ¿Finality?
	if MAX_PERCENTAGE == 0.7 {
		if _, ok := CHARACTERS[key]; ok {
			this.missingKeys[key] = 1
		} else {
			this.missingKeys[key] = this.missingKeys[key] + 1
		}
	}

	return value
}

func (this BuilderMstAscii) similarity(key string, char string) float64 {
	percentage := 0.0
	for i := range key {
		if key[i] == char[i] {
			percentage = percentage + 1;
		}
	}
	return percentage / float64(len(key))
}

func (this BuilderMstAscii) treatMissingKeys() {
	maxMissing := 2
	missingKeys := []string{}
	for key := range this.missingKeys {
		if this.missingKeys[key] > maxMissing {
			maxMissing = this.missingKeys[key]
			missingKeys = append(missingKeys, key)
		}
	}
	println("max missing: " + strconv.Itoa(maxMissing))
	for key := range this.missingKeys {
		println(key)
	}
}