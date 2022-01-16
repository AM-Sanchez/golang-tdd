package hello

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Guten Tag, "
const englishHelloDefaultSuffix = "World!!!"

func Hello(name, language string) string {

	if name == "" {
		name = englishHelloDefaultSuffix
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Human", ""))
}
