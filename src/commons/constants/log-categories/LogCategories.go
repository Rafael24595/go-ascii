package log_categories

type LogCategory string

const (
	INFO LogCategory= "INFO"
	WARNING LogCategory = "WARNING"
	ERROR LogCategory = "ERROR"
)

var dict = map[string]LogCategory{
    "INFO": INFO,
    "WARNING": WARNING,
	"ERROR": ERROR,
}

func LogCategoryList() (keys []string ){
	keys = make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	return keys
}