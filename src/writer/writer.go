package writer

import "fmt"

const defaultLanguage = "en"

type pair map[string]string

//Writer ...
type Writer struct {
	lang     string
	messages pair
	errors   pair
}

//GetMessage ...
func GetMessage(mType string) string {
	return Messages[mType]
}

//ShowMessage ...
func ShowMessage(mType string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(Messages[mType], a...))
}

//Init ...
func Init() {

}

//GetLanguage ...
func GetLanguage() string {
	return "en"
}

//SetLanguage ...
func SetLanguage(lang string) {

}
