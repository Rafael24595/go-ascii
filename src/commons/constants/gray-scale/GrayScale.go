package gray_scale

type GrayScale string

const (
	SOUP_CAHOS_HARD GrayScale = "SOUP_CAHOS_HARD"
	SOUP_CAHOS_SOFT GrayScale = "SOUP_CAHOS_SOFT"
	DEFAULT GrayScale = "DEFAULT"
)

var dict = map[string]GrayScale{
    "SOUP_CAHOS_HARD": "@MBHENR#KWXDFPQASUZbdehx*8Gm&04LOVYkpq5Tagns69owz$CIu23Jcfry%1v7l+it[] {}?j|()=~!-/<>\"^_';,:`. ",
    "SOUP_CAHOS_SOFT": "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'.",
	"DEFAULT": "@#S%?*+;:,.",
}

func GetGrayScale(code string) GrayScale  {
	return dict[string(code)]
}

func GetGrayScaleList() (keys []string ){
	keys = make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	return keys
}

func IsValidGrayScale(code string) bool {
	_, ok := dict[string(code)]
	return ok
}