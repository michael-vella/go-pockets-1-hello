package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. 'en'.")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"en": "Hello world",      // English
	"fr": "Bonjour le monde", // French
	"sp": "Hola mundo",       // Spanish
	"mt": "Hello dinja",      // Maltese
}

// greet returns a greeting to the world
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
