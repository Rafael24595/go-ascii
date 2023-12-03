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

type BuilderMstAscii struct {
	event ascii.QueueEvent
	images collection.ImagesCollection
	imgeScale scale.ImageScale
	grayScale gray_scale.GrayScale 
}

func NewBuilderMstAscii(event ascii.QueueEvent) (builder BuilderMstAscii, err error) {
	scaleHeight := event.GetHeight()
	scaleWidth := event.GetWidth()
	grayScale := event.GetGrayScale()

	images, err := encoder_decoder.DecodeByPath(event.GetPath())
	collection := collection.NewImagesCollection(images)
	scale := scale.NewImageScale(collection, scaleHeight, scaleWidth)
	builder = BuilderMstAscii{event: event, images:collection, imgeScale: scale, grayScale: grayScale}
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
	/*height := this.images.GetImageHeight()
	width := this.images.GetImageWidth()
	
	scaleX := this.imgeScale.GetScaleX()
	scaleY := this.imgeScale.GetScaleY()
	correction := this.getCorrection()*/

	grayscale := this.desaturateImage(position)

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

	number_of_characters_width := int(math.Floor(float64(grayscale.Rect.Dx()) / 8))
	number_of_characters_height := int(math.Floor(float64(grayscale.Rect.Dy()) / 16)) 
	if number_of_characters_width <= 1 && number_of_characters_height <= 1 {
		panic("Image is too small")
	} else {
		for x := 0; x < number_of_characters_width; x++ {
			line := []string{}
			for y := 0; y < number_of_characters_height; y++ {
				values := []string{}
				for i := 0; i < 8; i++ {
					for j := 0; j < 16; j++ {
						c := grayscale.GrayAt(int(x*8*i), int(y*16*j)).Y
						values = append(values, strconv.Itoa(int(c)))
					}
				}
				line = append(line, this.get_character(values))
			}
			print(strings.Join(line, "") + "\n")
		}
	}
	this.treat_missing_keys()

	return
}

func (this BuilderMstAscii) treat_missing_keys() {
	max_missing := 2
	missing_keys := []string{}
	for key := range thisMissingKeys {
		if thisMissingKeys[key] > max_missing {
			max_missing = thisMissingKeys[key]
			missing_keys = append(missing_keys, key)
		}
	}
	println("max missing: " + strconv.Itoa(max_missing))
	for key := range thisMissingKeys {
		println(key)
	}
}

var thisDict = map[string]string{
	"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111": "█" ,
	"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000": " " ,
	"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000110000001100000000000": "." ,
	"00011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000": "|" ,
	"11000011110000110110011001100110001001000011110000011000000110000001100000011000001111000010010001100110011001101100001111000011": "X" ,
	"00000000000000000000000000000000000000000000000000000000111111111111111100000000000000000000000000000000000000000000000000000000": "-" ,
	"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001111111111111111": "_" ,
	"11111111111111110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000111111111111111111": "]" ,
	"11111111111111111100000011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001111111111111111": "[" ,
	"00011000000110000001100000011000000110000001100000011000000110000001100000011000000111111111111111111111111111110000001100000011": "A" ,
	"11111111111111110000001100000011000000110000001111111111111111111111111111111110000001100000011000000110000001111111111111111111": "B",
	"00011111111111110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011": "C",
	//"11111111111111110000001100000011000000110000001100000011000000110000001100000011000000110000001100000011000000111111111111111111": "D",
	"11111111111111111100000011000000110000001100000011111111111111111100000011000000110000001100000011000000110000001100000011000000": "F",
	"11000000110000001110000001100000011100000011000000111000000110000001100000011100000011000000111000000110000001110000001100000011": "\\",
	"00000011000000110000011100000110000011100000110000011100000110000001100000111000001100000111000001100000111000001100000011000000": "/",
	"00000000011111110000000001111111000000000111111100000000011111110000000001111111000000000111111100000000011111110000000001111111": "▒",
}

var thisMissingKeys = map[string]int{

}

func (this BuilderMstAscii) get_character(values []string) string {
	char_map_key := strings.Join(values, "")
	max_percentage := 0.5
	max_percentage_char := " "

	for char, _ := range thisDict {
		percentage := this.get_percentage(char_map_key, char)
		if percentage > max_percentage {
			//max_percentage = percentage
			max_percentage_char = thisDict[char]
		}
	}

	if max_percentage == 0.7 {
		if _, ok := thisDict[char_map_key]; ok {
			thisMissingKeys[char_map_key] = 1
		} else {
			thisMissingKeys[char_map_key] = thisMissingKeys[char_map_key] + 1
		}
	}

	return max_percentage_char
}

func (this BuilderMstAscii) get_percentage(char_map_key string, char string) float64 {
	percentage := 0.0
	for i, _ := range char_map_key {
		x := i >= len(char_map_key)
		y := i >= len(char)
		if x || y {
			print("Oh!")
		}
		if char_map_key[i] == char[i] {
			percentage = percentage + 1;
		}
	}
	return percentage / float64(len(char_map_key))
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