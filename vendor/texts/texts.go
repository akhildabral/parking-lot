package texts

const defaultLanguage = "en"

type pair map[string]string

type Texts struct {
	lang     string
	messages pair
	errors   pair
}

func Init() {

}

func GetLanguage() string {
	return "en"
}

func SetLanguage(lang string) {

}
