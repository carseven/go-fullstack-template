package languages

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var languages = map[string]map[string]string{
	"en": {"HELLO": "Hello"},
	"es": {"HELLO": "Hola"},
}

func LoadLanguagesTranslations() {
	fmt.Println("[LANGUAGES] Init loading languages JSON files")
	// Load assets/languages json files and map to languages
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("[LANGUAGES][ERROR] Reading cwd directory" + err.Error() + ")")
		return
	}
	languagesPath := path.Join(cwd, "/assets/languages")

	// Check language file exist
	languagesFiles, err := os.ReadDir(languagesPath)
	if err != nil {
		fmt.Println("[LANGUAGES][ERROR] Reading /assets/languages directory")
		fmt.Println(err)
		return
	}

	// Read files and map to languages
	for _, file := range languagesFiles {
		if file.Type().IsRegular() {
			filePath := path.Join(languagesPath, file.Name())
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("[LANGUAGES][ERROR] Reading " + filePath)
				fmt.Println(err)
			}
			var rawJson map[string]interface{}
			json.Unmarshal([]byte(fileContent), &rawJson)
			fmt.Println(rawJson)
			// TODO: Map fileContent to languages hash map
			// Try https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map/21363587#21363587
			// Try https://www.sohamkamani.com/golang/json/
		}
	}
	fmt.Println("[LANGUAGES] TODO: Show successfully loaded languages")
}

func GetKeyTranslation(key string, userLanguages []string) string {
	var translation string
	for _, language := range userLanguages {
		translation = languages[language][key]
		if translation != "" {
			return translation
		}
	}

	// If no exist check key exist in the default language (en by default)
	translation = languages["en"][key]
	if translation != "" {
		return translation
	}

	// If no luck, just return the key string
	return key
}
